package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/linkrepository"
)

func Test_New(t *testing.T) {
	lr, err := linkrepository.New(&linkrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	_ = New(lr)
}
