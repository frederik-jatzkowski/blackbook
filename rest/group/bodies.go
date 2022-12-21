package group

type create struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type invite struct {
	GroupID   uint   `json:"groupId"`
	UserEmail string `json:"userEmail"`
	Message   string `json:"message"`
}

type acceptdecline struct {
	InvitationId uint `json:"invitationId"`
}

type invitationData struct {
	ID               uint   `json:"id"`
	Message          string `json:"message"`
	SenderFirstName  string `json:"senderFirstName"`
	SenderEmail      string `json:"senderEmail"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
}

type groupData struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type groupFeed struct {
	Invitations []invitationData `json:"invitations"`
	Groups      []groupData      `json:"groups"`
}
