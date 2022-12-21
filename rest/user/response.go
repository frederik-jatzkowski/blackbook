package user

import "github.com/frederik-jatzkowski/blackbook/database"

type responseUser struct {
	ID        uint   `json:"id"`
	Active    bool   `json:"active"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func newResponseUser(user *database.User) *responseUser {
	return &responseUser{
		ID:        user.ID,
		Active:    user.Active,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
