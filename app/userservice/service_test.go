package userservice

import (
	"testing"
)

func Test_New(t *testing.T) {
	_ = NewUserService(t, &Config{})
}
