package proxymailer

import (
	"log"
	"net/smtp"

	"github.com/gibbonco/titan/mailer"
	"github.com/jordan-wright/email"
)

type Mailer struct {
	ServerAddr   string
	Auth         smtp.Auth
	ProxyAddress string
}

func (m *Mailer) Send(email email.Email) *mailer.MailerError {
	if len(email.From) == 0 {
		email.From = "Arya App <info@aryaapp.co>"
	}

	// Replace the To address with our proxy address
	email.To = []string{m.ProxyAddress}

	err := email.Send(m.ServerAddr, m.Auth)
	if err != nil {
		log.Println(err)
		return &mailer.MailerError{}
	}
	return nil
}
