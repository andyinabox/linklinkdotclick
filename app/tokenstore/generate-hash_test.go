package tokenstore

import "testing"

func Test_GenerateHash(t *testing.T) {
	_, err := generateHash(32)
	if err != nil {
		t.Fatal(err.Error())
	}
}
