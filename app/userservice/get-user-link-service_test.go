package userservice

import (
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/test"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Test_GetUserLinkServiceInMemoryDb(t *testing.T) {
	email := "test@example.com"
	db := test.NewInMemoryDb(t)
	r := userrepository.New(db)
	s := New(&Config{
		DefaultUserEmail: email,
	}, r, tokenstore.New(db, &tokenstore.Config{}))

	user, err := s.EnsureDefaultUser()
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = s.GetUserLinkService(user)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_GetUserLinkServiceFsDb(t *testing.T) {

	usersDb := "../../test/db/test.db"
	linksDbDir := path.Join(path.Dir(usersDb), "usr")

	err := os.MkdirAll(path.Dir(linksDbDir), fs.ModePerm)
	if err != nil {
		t.Fatal(err.Error())
	}

	db, err := gorm.Open(sqlite.Open(usersDb), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}

	r := userrepository.New(db)

	email := "test@example.com"
	s := New(&Config{
		UserDbPath:       linksDbDir,
		DefaultUserEmail: email,
	}, r, tokenstore.New(db, &tokenstore.Config{}))

	user, err := s.EnsureDefaultUser()
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = s.GetUserLinkService(user)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Cleanup(func() {
		os.RemoveAll(path.Dir(usersDb))
	})
}
