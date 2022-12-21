package user

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type create struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

type activate struct {
	ActivationCode string `json:"activationCode"`
}

type update struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type delete struct {
	Email string `json:"email"`
}
