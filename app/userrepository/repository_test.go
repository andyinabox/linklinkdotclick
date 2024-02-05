package userrepository

<<<<<<< HEAD
import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_New(t *testing.T) {
	db := test.NewInMemoryDb(t)
	_ = New(db)
=======
import "testing"

func Test_New(t *testing.T) {
	_, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}
>>>>>>> main
}
