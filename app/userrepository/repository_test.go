package userrepository

import "testing"

func Test_New(t *testing.T) {
	_, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}
}
