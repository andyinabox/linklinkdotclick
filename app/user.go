package app

import "time"

type User struct {
	// gorm fields
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// domain fields
	Email string `json:"email" gorm:"uniqueIndex"`
}

type UserRepository interface {
	FetchUser(id uint) (*User, error)
	FetchUserByEmail(email string) (*User, error)
	CreateUser(User) (*User, error)
	UpsertUser(User) (*User, error)
}

type UserService interface {
	CreateUser(email string) (*User, error)
	FetchUser(id uint) (*User, error)
	FetchOrCreateUserByEmail(email string) (*User, error)
	EnsureDefaultUser() (*User, error)
	GetUserLinkService(*User) (LinkService, error)
	GetLoginHashForUser(*User) (string, error)
	GetUserFromLoginHash(string) (*User, error)
}
