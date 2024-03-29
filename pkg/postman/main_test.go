package postman

import (
	"net/http"
	"net/mail"
	"testing"
)

func Test_Send(t *testing.T) {

	// test if mailpit is running
	resp, _ := http.Get("http://127.0.0.1:8025/")
	if resp == nil || resp.StatusCode != http.StatusOK {
		t.Skip("mailpit doesn't seem to be running, skipping mail test")
		return
	}

	email := &Email{
		From:    mail.Address{"Jack Handy", "jack@example.com"},
		To:      mail.Address{"Jill Dandy", "jill@example.com"},
		Subject: "Howdy",
		Body:    "How are you doing??",
	}

	err := Send(email, "127.0.0.1:1025")
	if err != nil {
		t.Fatal(err.Error())
	}
}
