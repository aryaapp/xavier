package mailer

import (
	"bytes"
	"text/template"

	"github.com/jordan-wright/email"
)

type Mailer interface {
	Send(email.Email) *MailerError
}

type MailerError struct {
}

func (m *MailerError) Error() string {
	return "Mail shit"
}

func NewVerifyEmail(emailAddr, signupUrl string) email.Email {
	buf := new(bytes.Buffer)
	templ := template.Must(template.New("verify_email").ParseFiles(
		"templates/mailer/base.html",
		"templates/mailer/welcome_email.html"))
	templ.ExecuteTemplate(buf, "arya", map[string]interface{}{
		"Username":  emailAddr,
		"SignupUrl": signupUrl,
	})

	return email.Email{
		To:      []string{emailAddr},
		Subject: "Please verify your email for Arya App",
		HTML:    buf.Bytes(),
	}
}

func NewWelcomeEmail(emailAddr) email.Email {
	buf := new(bytes.Buffer)
	templ := template.Must(template.New("welcome_email").ParseFiles(
		"templates/mailer/base.html",
		"templates/mailer/welcome_email.html"))
	templ.ExecuteTemplate(buf, "arya", map[string]interface{}{
		"Username": emailAddr,
	})

	return email.Email{
		To:      []string{emailAddr},
		Subject: "Welcome to Arya App",
		HTML:    buf.Bytes(),
	}
}
