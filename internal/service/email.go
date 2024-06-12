package service

import (
	"github.com/CeoFred/gin-boilerplate/constants"
	"github.com/CeoFred/gin-boilerplate/sendgrid"
)

type EmailServicer interface {
	SendForgotPasswordEmail(name, email string)
}

type EmailService struct {
	Client *sendgrid.Client
}

var (
	constant = constants.New()
)

func NewEmailService() EmailServicer {
	client := sendgrid.NewClient(constant.SendGridApiKey)
	return &EmailService{Client: client}
}

func (s *EmailService) Send(name, email, subject, content string) error {
	to := sendgrid.EmailAddress{
		Name:  name,
		Email: email,
	}

	return s.Client.Send(&to, constant.SendFromEmail, constant.SendFromName, subject, content)

}

func (s *EmailService) SendForgotPasswordEmail(name, email string) {
	//TODO: send otp email to user
}

// Sends account verification email
