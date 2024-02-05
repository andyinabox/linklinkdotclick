package linkservice

import (
	"testing"
)

func Test_CreateLink(t *testing.T) {
	ls := NewLinkService(t)
	_, err := ls.CreateLink("https://www.w3.org/blog/")
	if err != nil {
		t.Fatal(err.Error())
	}
}
