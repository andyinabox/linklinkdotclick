package userservice

import (
<<<<<<< HEAD
	"io/fs"
=======
>>>>>>> main
	"os"
	"path"
	"testing"

<<<<<<< HEAD
	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Test_GetUserLinkServiceInMemoryDb(t *testing.T) {
	email := "test@example.com"
	s := NewUserService(t, &Config{
		DefaultUserEmail: email,
	})

=======
	"github.com/andyinabox/linkydink/app/userrepository"
)

func Test_GetUserLinkServiceInMemoryDb(t *testing.T) {
	r, err := userrepository.New(&userrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	email := "test@example.com"
	s := New(&Config{
		DefaultUserEmail: email,
	}, r)
>>>>>>> main
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

<<<<<<< HEAD
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
	config := &Config{
		UserDbPath:       linksDbDir,
		DefaultUserEmail: email,
	}
	ts := tokenstore.New(db, &tokenstore.Config{})
	s := New(r, ts, config)

=======
	r, err := userrepository.New(&userrepository.Config{
		DbFile: usersDb,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	email := "test@example.com"
	s := New(&Config{
		UserDbPath:       linksDbDir,
		DefaultUserEmail: email,
	}, r)
>>>>>>> main
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
