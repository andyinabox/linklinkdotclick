package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/linkrepository"
)

func Test_New(t *testing.T) {
	lr := linkrepository.New(&linkrepository.Config{
		DbFile: ":memory:",
	})
	_ = New(lr)
}
