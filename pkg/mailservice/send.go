package mailservice

import (
	"bytes"
	"net/smtp"
)

// https://gist.github.com/jniltinho/d90034994f29d7d25e59c9e0fe5548d2
func (s *Service) Send(email *Email) (err error) {

	client, err := smtp.Dial(s.conf.SmtpAddr)
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
	err = s.tmpl.Execute(buf, email)
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
