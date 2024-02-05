package userservice

import (
	"testing"
)

func Test_CreateUser(t *testing.T) {
	s := NewUserService(t, &Config{})
	invalidEmail := "example.com"
	_, err := s.CreateUser(invalidEmail)
	if err == nil {
		t.Errorf("Expected email validation error for '%s', got none", invalidEmail)
	}

	validEmail := "test@example.com"
	_, err = s.CreateUser(validEmail)
	if err != nil {
		t.Fatal(err.Error())
	}

}
