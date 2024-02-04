package mailservice

import "testing"

func Test_New(t *testing.T) {
	_ = New(&Config{})
}
