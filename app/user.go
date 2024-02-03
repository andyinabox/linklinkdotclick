package app

import "time"

type User struct {
	// gorm fields
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// domain fields
	Email string
}

type UserRepository interface {
	FetchUser(id uint) (*User, error)
	CreateUser(User) (*User, error)
}

type UserService interface {
	FetchUser(id uint) (*User, error)
	CreateUser(email string) (*User, error)
	GetUserLinkService(User) (LinkService, error)
}
