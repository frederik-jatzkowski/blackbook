package group

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frederik-jatzkowski/blackbook/database"
	"github.com/frederik-jatzkowski/blackbook/user"
	"github.com/frederik-jatzkowski/blackbook/util"
	"gorm.io/gorm"
)

type Service struct {
	db   *gorm.DB
	user *user.Service
}

func NewService(db *gorm.DB, user *user.Service) (*Service, error) {
	if db == nil {
		return nil, fmt.Errorf("db cannot be nil")
	}

	if user == nil {
		return nil, fmt.Errorf("user service cannot be nil")
	}

	return &Service{
		db:   db,
		user: user,
	}, nil
}

func (service Service) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var (
		body   create
		err    error
		user   *database.User
		group  database.Group
		result *gorm.DB
	)

	body, err = util.ParseBody[create](w, r, "POST")
	if err != nil {
		return
	}

	// get user
	user = service.user.GetActiveUserFromRequest(r)
	if user == nil {
		util.WriteResponse(w, nil, "Bitte anmelden.")

		return
	}

	group = database.Group{
		Name:        body.Name,
		Description: body.Description,
		Users:       []database.User{*user},
	}

	result = service.db.Create(&group)
	if result.Error != nil {
		util.WriteResponse(w, user, "Gruppe konnte nicht erstellt werden.")

		return
	}

	util.WriteResponse(w, user)
}

func (service Service) HandleInvite(w http.ResponseWriter, r *http.Request) {
	var (
		body       invite
		err        error
		sender     *database.User
		receiver   *database.User
		group      *database.Group
		invitation database.Invitation
		result     *gorm.DB
		count      int64
	)

	body, err = util.ParseBody[invite](w, r, "POST")
	if err != nil {
		return
	}

	// get inviting user
	sender = service.user.GetActiveUserFromRequest(r)
	if sender == nil {
		util.WriteResponse(w, sender, "Bitte anmelden.")

		return
	}

	// find group
	for _, group2 := range sender.Groups {
		if group2.ID == body.GroupID {
			group = &group2

			break
		}
	}

	if group == nil {
		util.WriteResponse(w, sender, "Kein Mitglied dieser Gruppe.")

		return
	}

	// get invited user
	result = service.db.Where("email = ?", body.UserEmail).Preload("Groups").Take(&receiver)
	if result.Error != nil {
		util.WriteResponse(w, sender, "Eingeladener Nutzer existiert nicht.")

		return
	}

	// check, if user is in group
	for _, group2 := range receiver.Groups {
		if group.ID == group2.ID {
			util.WriteResponse(w, sender, "Nutzer ist bereits Teil der Gruppe.")

			return
		}
	}

	// add new invitation
	invitation = database.Invitation{
		ReceiverID: receiver.ID,
		GroupID:    group.ID,
	}

	// find invitation
	result = service.db.Where(&invitation).Count(&count)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error while counting invitations: %s", err)

		return
	}

	if count != 0 {
		util.WriteResponse(w, sender, "Nutzer wurde bereits eingeladen.")

		return
	}

	invitation.SenderID = sender.ID
	invitation.Message = body.Message

	result = service.db.Create(&invitation)
	if result.Error != nil {
		util.WriteResponse(w, sender, "Einladung konnte nicht erstellt werden.")

		return
	}

	util.WriteResponse(w, sender)
}

func (service Service) HandleAccept(w http.ResponseWriter, r *http.Request) {
	var (
		body       acceptdecline
		err        error
		user       *database.User
		invitation database.Invitation
		result     *gorm.DB
	)

	body, err = util.ParseBody[acceptdecline](w, r, "POST")
	if err != nil {
		return
	}

	// get invited user
	user = service.user.GetActiveUserFromRequest(r)
	if user == nil {
		util.WriteResponse(w, nil, "Bitte anmelden.")

		return
	}

	// get invitation
	result = service.db.Where("receiver_id = ?", user.ID).Preload("Group").Take(&invitation, body.InvitationId)
	if result.Error != nil {
		util.WriteResponse(w, user, "Einladung existiert nicht.")

		return
	}

	// add user to group
	err = service.db.Model(&user).Association("Groups").Append(&invitation.Group)
	if err != nil {
		util.WriteResponse(w, user, "Konnte nicht zur Gruppe hinzugefügt werden.")
		log.Printf("could not add user to group: %s", err)

		return
	}

	// delete invitation
	result = service.db.Delete(&invitation)
	if result.Error != nil {
		log.Printf("could not delete invitation with id %d: %s", invitation.ID, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	util.WriteResponse(w, user)
}

func (service Service) HandleDecline(w http.ResponseWriter, r *http.Request) {
	var (
		body       acceptdecline
		err        error
		user       *database.User
		invitation database.Invitation
		result     *gorm.DB
	)

	body, err = util.ParseBody[acceptdecline](w, r, "POST")
	if err != nil {
		return
	}

	// get invited user
	user = service.user.GetActiveUserFromRequest(r)
	if user == nil {
		util.WriteResponse(w, nil, "Bitte anmelden.")

		return
	}

	// get invitation
	result = service.db.Where("receiver_id = ?", user.ID).Preload("Group").Take(&invitation, body.InvitationId)
	if result.Error != nil {
		util.WriteResponse(w, user, "Einladung existiert nicht.")

		return
	}

	// delete invitation
	result = service.db.Delete(&invitation)
	if result.Error != nil {
		util.WriteResponse(w, user, "Einladung konnte nicht gelöscht werden.")
		log.Printf("could not delete invitation with id %d: %s", invitation.ID, err)

		return
	}

	util.WriteResponse(w, user)
}

func (service Service) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (service Service) HandleLeave(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
