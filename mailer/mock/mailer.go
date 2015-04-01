package mockmailer

import (
	"github.com/gibbonco/titan/mailer"
	"github.com/jordan-wright/email"
)

type Mailer struct{}

var SendEmails = []*email.Email{}

func (m *Mailer) Send(email email.Email) *mailer.MailerError {
	if len(email.From) == 0 {
		email.From = "Arya App <info@aryaapp.co>"
	}
	SendEmails = append(SendEmails, &email)
	return nil
}
