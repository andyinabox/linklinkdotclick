package userrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/test"
)

func Test_UpsertUser(t *testing.T) {
	db := test.NewInMemoryDb(t)
	r := New(db)

	expectedEmail := "test2@example.com"

	// upsert user
	user, err := r.UpsertUser(app.User{
		ID:    1,
		Email: "test1@example.com",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	// update user
	user.Email = expectedEmail

	updatedUser, err := r.UpsertUser(*user)
	if err != nil {
		t.Fatal(err.Error())
	}

	if updatedUser.Email != expectedEmail {
		t.Errorf("expected updated user email to be %s, got %s", expectedEmail, updatedUser.Email)
	}

}
