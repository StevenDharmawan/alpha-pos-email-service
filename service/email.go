package service

import (
	"email-service/config"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

type EmailService interface {
	SendEmail(userEmail string) error
}

type EmailServiceImpl struct {
	*config.EmailConfig
}

func NewEmailService(emailConfig *config.EmailConfig) *EmailServiceImpl {
	return &EmailServiceImpl{EmailConfig: emailConfig}
}

func (service *EmailServiceImpl) SendEmail(subject string, userEmail string, message string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", service.EmailConfig.Name)
	mailer.SetHeader("To", userEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)
	fmt.Println("test")
	dialer := gomail.NewDialer(
		service.EmailConfig.Host,
		service.EmailConfig.Port,
		service.EmailConfig.Email,
		service.EmailConfig.Password,
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Success")
}
