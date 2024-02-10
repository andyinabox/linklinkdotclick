package postman

import (
	"bytes"
	"net/mail"
	"net/smtp"
	"strings"
	"text/template"
)

const MimeHtml = `MIME-version: 1.0;
Content-Type: text/html; charset="UTF-8"`

var emailTemplateBody = strings.TrimSpace(`
From: {{ .From }}
To: {{ .To }}
Subject: {{ .Subject }}
{{ .Mime }}

{{ .Body }}
`)

var emailTemplate *template.Template

type Email struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Mime    string
	Body    string
}

func init() {
	emailTemplate = template.New("email.txt.tmpl")
	emailTemplate = template.Must(emailTemplate.Parse(emailTemplateBody))
}

// https://gist.github.com/jniltinho/d90034994f29d7d25e59c9e0fe5548d2
func Send(email *Email, smtpAddr string) (err error) {

	client, err := smtp.Dial(smtpAddr)
	if err != nil {
		return
	}
	defer client.Close()

	// send MAIL command to server
	err = client.Mail(email.From.String())
	if err != nil {
		return
	}

	// RCPT command
	err = client.Rcpt(email.To.String())
	if err != nil {
		return
	}

	w, err := client.Data()
	if err != nil {
		return
	}

	buf := &bytes.Buffer{}
	err = emailTemplate.Execute(buf, email)
	if err != nil {
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		return
	}

	err = w.Close()
	if err != nil {
		return
	}

	err = client.Quit()

	return
}
