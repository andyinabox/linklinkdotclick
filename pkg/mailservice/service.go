package mailservice

import (
	"strings"
	"text/template"
)

const MimeHtml = `MIME-version: 1.0;
Content-Type: text/html; charset="UTF-8"`

var emailTemplate = strings.TrimSpace(`
From: {{ .From }}
To: {{ .To }}
Subject: {{ .Subject }}
{{ .Mime }}

{{ .Body }}
`)

type Config struct {
	SmtpAddr string
}

type Service struct {
	conf *Config
	tmpl *template.Template
}

func New(conf *Config) *Service {

	tmpl := template.New("email.txt.tmpl")
	tmpl = template.Must(tmpl.Parse(emailTemplate))

	return &Service{conf, tmpl}
}
