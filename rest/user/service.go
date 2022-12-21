package user

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/frederik-jatzkowski/blackbook/database"
	"github.com/frederik-jatzkowski/blackbook/mail"
	"github.com/frederik-jatzkowski/blackbook/util"
	"gorm.io/gorm"
)

type Service struct {
	pepper string
	db     *gorm.DB
	mailer *mail.Service
}

func NewService(db *gorm.DB, mailer *mail.Service) (*Service, error) {
	pepper := os.Getenv("PEPPER")

	if len(pepper) < 16 {
		return nil, fmt.Errorf("pepper must be at least 16 characters long but was %d", len(pepper))
	}

	return &Service{
		pepper: pepper,
		db:     db,
		mailer: mailer,
	}, nil
}

func (service *Service) getUserFromRequest(r *http.Request) *database.User {
	var (
		sessionCookie *http.Cookie
		session       string
		id            uint
		user          database.User
		err           error
	)

	// find session cookie
	sessionCookie = findSessionCookie(r)
	if sessionCookie == nil {
		return nil
	}

	// extract info from cookie
	id, session, err = parseSessionCookie(sessionCookie)
	if err != nil {
		return nil
	}

	// find user
	if service.db.Model(&database.User{}).Preload("Groups").Take(&user, id).Error != nil {
		return nil
	}

	// check session end time
	if user.SessionExpiration.Before(time.Now()) {
		return nil
	}

	// check session token
	if subtle.ConstantTimeCompare([]byte(session), []byte(user.Session)) == 0 {
		return nil
	}

	return &user
}

func (service *Service) GetActiveUserFromRequest(w http.ResponseWriter, r *http.Request) *database.User {
	user := service.getUserFromRequest(r)
	if user == nil {
		util.WriteResponse[*responseUser](w, nil, nil, "Bitte anmelden.")

		return nil
	}

	// activity check
	if !user.Active {
		util.WriteResponse(w, newResponseUser(user), nil, "Aktivierung nicht abgeschlossen")

		return nil
	}

	return user
}

func (service *Service) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var (
		body login
		user *database.User
		err  error
	)

	// parse body
	body, err = util.ParseBody[login](w, r, "POST")
	if err != nil {
		return
	}

	// retrieve user with email
	result := service.db.Where("email = LOWER(?)", body.Email).Take(&user)
	if result.Error != nil {
		util.WriteResponse[*responseUser](w, nil, nil, "Ungültige Anmeldedaten.")

		return
	}

	// compare hashes
	if !authenticate(user.Hash, body.Password, user.Salt, service.pepper) {
		util.WriteResponse[*responseUser](w, nil, nil, "Ungültige Anmeldedaten.")

		return
	}

	// create session
	user.Session = randBytes(64)
	user.SessionExpiration = time.Now().Add(time.Hour * 24 * 30)

	// set session cookie
	http.SetCookie(w, newSessionCookie(user.ID, user.Session))

	// persist session to db
	if service.db.Save(user).Error != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	util.WriteResponse(w, newResponseUser(user), util.PointerFor("Angemeldet."))
}

func (service *Service) HandleLogout(w http.ResponseWriter, r *http.Request) {
	user := service.getUserFromRequest(r)
	if user == nil {
		return
	}

	// create session
	user.Session = ""
	user.SessionExpiration = time.Time{}

	// persist session to db
	if service.db.Save(user).Error != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	util.WriteResponse[*responseUser](w, nil, util.PointerFor("Abgemeldet."))
}

func (service *Service) HandleSessionCheck(w http.ResponseWriter, r *http.Request) {
	user := service.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	// send success message
	util.WriteResponse(w, newResponseUser(user), nil)
}

func (service *Service) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var (
		body   create
		user   *database.User
		result *gorm.DB
		errs   []string
		err    error
	)

	body, err = util.ParseBody[create](w, r, "POST")
	if err != nil {
		return
	}

	// validate user input
	if body.Email == "" {
		errs = append(errs, "Bitte Emailadresse angeben.")
	}

	if body.FirstName == "" {
		errs = append(errs, "Bitte Vornamen angeben.")
	}

	if body.LastName == "" {
		errs = append(errs, "Bitte Nachnamen angeben.")
	}

	if len(body.Password) < 16 {
		errs = append(errs, "Das Passwort muss mindestens 16 Zeichen lang sein.")
	}

	if body.Password != body.PasswordRepeat {
		errs = append(errs, "Die beiden Passwörter stimmen nicht überein.")
	}

	if len(errs) > 0 {
		util.WriteResponse[*responseUser](w, nil, nil, errs...)

		return
	}

	// build user
	user = &database.User{
		FirstName:         body.FirstName,
		LastName:          body.LastName,
		Email:             body.Email,
		Salt:              randBytes(16),
		Active:            false,
		ActivationCode:    randBytes(15),
		Session:           randBytes(60),
		SessionExpiration: time.Now().Add(time.Hour * 24 * 30),
	}
	user.Hash = hashPassword(body.Password, user.Salt, service.pepper)

	// check if user exists
	if service.db.Where("email = LOWER(?)", user.Email).Take(&database.User{}).Error == nil {
		util.WriteResponse[*responseUser](w, nil, nil, "Ein Nutzer mit dieser Emailadresse existiert bereits.")

		return
	}

	// send activation code to user
	err = service.mailer.SendActivationCode(user)
	if err != nil {
		util.WriteResponse[*responseUser](w, nil, nil, "Fehler beim Zustellen des Aktivierungscodes.")
		log.Printf("could not send activation code: %s", err)

		return
	}

	// persist user
	result = service.db.Create(&user)
	if result.Error != nil {
		util.WriteResponse[*responseUser](w, nil, nil, "Fehler beim Abschluss des Registrierungsprozesses.")
		log.Printf("could not persist user: %s", result.Error)

		return
	}

	// set session cookie
	http.SetCookie(w, newSessionCookie(user.ID, user.Session))

	util.WriteResponse(w, newResponseUser(user), util.PointerFor("Konto erstellt."))
}

func (service *Service) HandleActivate(w http.ResponseWriter, r *http.Request) {
	var (
		body   activate
		user   *database.User
		result *gorm.DB
		err    error
	)

	body, err = util.ParseBody[activate](w, r, "POST")
	if err != nil {
		return
	}

	user = service.getUserFromRequest(r)
	if user == nil {
		util.WriteResponse[*responseUser](w, nil, nil, "Bitte anmelden.")

		return
	}

	if user.ActivationCode != body.ActivationCode {
		util.WriteResponse(w, newResponseUser(user), nil, "Fehlerhafter Aktivierungscode.")

		return
	}

	user.Active = true

	result = service.db.Save(&user)
	if result.Error != nil {
		util.WriteResponse(w, newResponseUser(user), nil, "Aktivierung fehlgeschlagen.")
		log.Printf("could not persist user: %s", result.Error)

		return
	}

	util.WriteResponse(w, newResponseUser(user), util.PointerFor("Konto aktiviert."))
}

func (service *Service) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		body   update
		user   *database.User
		result *gorm.DB
		err    error
	)

	body, err = util.ParseBody[update](w, r, "POST")
	if err != nil {
		return
	}

	user = service.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName

	result = service.db.Save(&user)
	if result.Error != nil {
		util.WriteResponse(w, newResponseUser(user), nil, "Änderungen fehlgeschlagen.")
		log.Printf("could not change user: %s", result.Error)

		return
	}

	util.WriteResponse(w, newResponseUser(user), util.PointerFor("Kontodaten geändert."))
}

func (service *Service) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var (
		body   login
		user   *database.User
		result *gorm.DB
		err    error
	)

	body, err = util.ParseBody[login](w, r, "POST")
	if err != nil {
		return
	}

	user = service.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	if user.Email != body.Email || !authenticate(user.Hash, body.Password, user.Salt, service.pepper) {
		util.WriteResponse(w, newResponseUser(user), nil, "Zugangsdaten nicht korrekt.")

		return
	}

	result = service.db.Select(
		"Groups",
		"SentInvitations",
		"ReceivedInvitations",
	).Delete(&user)
	if result.Error != nil {
		util.WriteResponse(w, newResponseUser(user), nil, "Löschung fehlgeschlagen.")
		log.Printf("could not change user: %s", result.Error)

		return
	}

	util.WriteResponse[*responseUser](w, nil, util.PointerFor("Konto gelöscht."))
}
