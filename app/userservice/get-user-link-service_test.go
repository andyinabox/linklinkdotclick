package userservice

import (
	"os"
	"path"
	"testing"

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
