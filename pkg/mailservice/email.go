package mailservice

import (
	"net/mail"
)

type Email struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Mime    string
	Body    string
}
