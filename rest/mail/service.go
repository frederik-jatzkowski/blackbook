package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/frederik-jatzkowski/blackbook/database"
)

type Service struct {
	plainAuth smtp.Auth
	from      string
	host      string
	port      string
	frontend  string
}

func NewService() (*Service, error) {
	var (
		frontend = os.Getenv("APP_NAME")
		identity = os.Getenv("SMTP_IDENTITY")
		username = os.Getenv("SMTP_USERNAME")
		password = os.Getenv("SMTP_PASSWORD")
		host     = os.Getenv("SMTP_HOST")
		port     = os.Getenv("SMTP_PORT")
		from     = username
	)

	if username == "" {
		return nil, fmt.Errorf("missing env variable SMTP_USERNAME")
	}

	if password == "" {
		return nil, fmt.Errorf("missing env variable SMTP_PASSWORD")
	}

	if host == "" {
		return nil, fmt.Errorf("missing env variable SMTP_HOST")
	}

	if port == "" {
		return nil, fmt.Errorf("missing env variable SMTP_PORT")
	}

	if frontend == "" {
		return nil, fmt.Errorf("missing env variable APP_NAME")
	}

	return &Service{
		plainAuth: smtp.PlainAuth(identity, username, password, host),
		from:      from,
		host:      host,
		port:      port,
		frontend:  frontend,
	}, nil
}

func (service *Service) SendActivationCode(user *database.User) error {
	if user.ActivationCode == "" {
		return fmt.Errorf("user does not have an activation code set")
	}

	message := []byte(fmt.Sprintf(
		"%s"+
			"Hallo %s,\r\n"+
			"\r\n"+
			"vielen Dank für deine Registrierung bei %s.\r\n"+
			"\r\n"+
			"Hier ist dein Aktivierungscode: %s\r\n"+
			"\r\n"+
			"Liebe Grüße\r\n",
		buildHeaders(service.from, user.Email, "Aktivierungscode", service.frontend),
		user.FirstName,
		service.frontend,
		user.ActivationCode,
	))

	return smtp.SendMail(service.host+":"+service.port, service.plainAuth, service.from, []string{user.Email}, []byte(message))
}
