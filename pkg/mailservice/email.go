package mailservice

import (
	"fmt"
	"net/mail"
)

const emailTemplate = `From: %s
To: %s
Subject: %s

%s
`

type Email struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Body    string
}

func (e *Email) String() string {
	return fmt.Sprintf(emailTemplate, e.From.String(), e.To.String(), e.Subject, e.Body)
}

func (e *Email) Bytes() []byte {
	return []byte(e.String())
}
