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
