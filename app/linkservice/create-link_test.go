package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/linkrepository"
)

func Test_CreateLink(t *testing.T) {
	lr, err := linkrepository.New(&linkrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	ls := New(lr)
	_, err = ls.CreateLink("https://www.w3.org/blog/")
	if err != nil {
		t.Fatal(err.Error())
	}
}
