package userrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
)

func Test_CreateUser(t *testing.T) {
	r, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}
	user := app.User{
		Email: "test@example.com",
	}
	createdUser, err := r.CreateUser(user)
	if err != nil {
		t.Fatal(err.Error())
	}
	if createdUser.ID == 0 {
		t.Fatal("expected id to be non-zero value")
	}
}
