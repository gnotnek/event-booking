package email

import (
	"event-booking/internal/config"
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	config *config.Smtp
}

func NewEmailService(config *config.Smtp) *EmailService {
	return &EmailService{
		config: config,
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.config.FromEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.config.SmtpHost, e.config.SmtpPort, e.config.Username, e.config.Password)

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

func (e *EmailService) SendVerificationEmail(to, code string) error {
	subject := "Email Verification Code"
	body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>Email Verification</title>
        </head>
        <body>
            <h1>Email Verification Code</h1>
            <p>Your verification code is: <strong>%s</strong></p>
            <p>Please use it within 1 hour.</p>
            <p>If you did not request this code, please ignore this email.</p>
        </body>
        </html>
    `, code)
	return e.SendEmail(to, subject, body)
}
