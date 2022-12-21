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

func (service Service) finishResponse(w http.ResponseWriter, user *database.User, success *string, errors ...string) {
	var (
		groups      []database.Group
		invitations []database.Invitation
		feed        groupFeed
		err         error
	)

	err = service.db.Order("id asc").Model(&user).Association("Groups").Find(&groups)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error while fetching groups: %s", err)

		return
	}

	err = service.db.Order("id asc").Model(&user).Preload("Sender").Preload("Group").
		Association("ReceivedInvitations").Find(&invitations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error while fetching invitations: %s", err)

		return
	}

	for _, group := range groups {
		feed.Groups = append(feed.Groups, groupData{
			ID:          group.ID,
			Name:        group.Name,
			Description: group.Description,
		})
	}

	for _, invitation := range invitations {
		feed.Invitations = append(feed.Invitations, invitationData{
			ID:               invitation.ID,
			Message:          invitation.Message,
			SenderFirstName:  invitation.Sender.FirstName,
			SenderEmail:      invitation.Sender.Email,
			GroupName:        invitation.Group.Name,
			GroupDescription: invitation.Group.Description,
		})
	}

	util.WriteResponse(w, feed, success, errors...)
}

func (service Service) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	var (
		user *database.User
	)

	// get user
	user = service.user.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	service.finishResponse(w, user, nil)
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
	user = service.user.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	group = database.Group{
		Name:        body.Name,
		Description: body.Description,
		Users:       []database.User{*user},
	}

	result = service.db.Create(&group)
	if result.Error != nil {
		service.finishResponse(w, user, nil, "Gruppe konnte nicht erstellt werden.")

		return
	}

	service.finishResponse(w, user, util.PointerFor("Gruppe erstellt."))
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
	sender = service.user.GetActiveUserFromRequest(w, r)
	if sender == nil {
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
		service.finishResponse(w, sender, nil, "Kein Mitglied dieser Gruppe.")

		return
	}

	// get invited user
	result = service.db.Where("email = ?", body.UserEmail).Preload("Groups").Take(&receiver)
	if result.Error != nil {
		service.finishResponse(w, sender, nil, "Eingeladener Nutzer existiert nicht.")

		return
	}

	// check, if user is in group
	for _, group2 := range receiver.Groups {
		if group.ID == group2.ID {
			service.finishResponse(w, sender, nil, "Nutzer ist bereits Teil der Gruppe.")

			return
		}
	}

	// add new invitation
	invitation = database.Invitation{
		ReceiverID: receiver.ID,
		GroupID:    group.ID,
	}

	// find invitation
	result = service.db.Model(&invitation).Where(&invitation).Count(&count)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error while counting invitations: %s", result.Error)

		return
	}

	if count != 0 {
		service.finishResponse(w, sender, nil, "Nutzer wurde bereits eingeladen.")

		return
	}

	invitation.SenderID = sender.ID
	invitation.Message = body.Message

	result = service.db.Create(&invitation)
	if result.Error != nil {
		service.finishResponse(w, sender, nil, "Einladung konnte nicht erstellt werden.")

		return
	}

	service.finishResponse(w, sender, util.PointerFor("Einladung abgeschickt."))
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
	user = service.user.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	// get invitation
	result = service.db.Where("receiver_id = ?", user.ID).Preload("Group").Take(&invitation, body.InvitationId)
	if result.Error != nil {
		service.finishResponse(w, user, nil, "Einladung existiert nicht.")

		return
	}

	// add user to group
	err = service.db.Model(&user).Association("Groups").Append(&invitation.Group)
	if err != nil {
		service.finishResponse(w, user, nil, "Konnte nicht zur Gruppe hinzugefügt werden.")
		log.Printf("could not add user to group: %s", err)

		return
	}

	// delete invitation
	result = service.db.Delete(&invitation)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not delete invitation with id %d: %s", invitation.ID, err)

		return
	}

	service.finishResponse(w, user, util.PointerFor("Gruppe beigetreten."))
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
	user = service.user.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	// get invitation
	result = service.db.Where("receiver_id = ?", user.ID).Preload("Group").Take(&invitation, body.InvitationId)
	if result.Error != nil {
		service.finishResponse(w, user, nil, "Einladung existiert nicht.")

		return
	}

	// delete invitation
	result = service.db.Delete(&invitation)
	if result.Error != nil {
		service.finishResponse(w, user, nil, "Einladung konnte nicht gelöscht werden.")
		log.Printf("could not delete invitation with id %d: %s", invitation.ID, result.Error)

		return
	}

	service.finishResponse(w, user, util.PointerFor("Einladung abgelehnt."))
}

func (service Service) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		body   groupData
		err    error
		user   *database.User
		result *gorm.DB
	)

	body, err = util.ParseBody[groupData](w, r, "POST")
	if err != nil {
		return
	}

	user = service.user.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	result = service.db.Model(&database.Group{ID: body.ID}).Updates(&database.Group{
		Name:        body.Name,
		Description: body.Description,
	})
	if result.Error != nil {
		service.finishResponse(w, user, nil, "Gruppe konnte nicht geändert werden.")
		log.Printf("could not update group with id %d: %s", body.ID, result.Error)

		return
	}

	service.finishResponse(w, user, util.PointerFor("Gruppe geändert."))
}

func (service Service) HandleLeave(w http.ResponseWriter, r *http.Request) {
	var (
		body groupData
		err  error
		user *database.User
	)

	body, err = util.ParseBody[groupData](w, r, "POST")
	if err != nil {
		return
	}

	user = service.user.GetActiveUserFromRequest(w, r)
	if user == nil {
		return
	}

	err = service.db.Model(&user).Association("Groups").Delete(&database.Group{ID: body.ID})
	if err != nil {
		service.finishResponse(w, user, nil, "Gruppe konnte nicht verlassen werden.")
		log.Printf("could not leave group with id %d: %s", body.ID, err)

		return
	}

	service.finishResponse(w, user, util.PointerFor("Gruppe verlassen."))
}
