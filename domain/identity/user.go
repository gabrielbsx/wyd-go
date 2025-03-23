package identity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Username string
	Password string
	Email    string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewUser(name, username, email, passwordHash string) (*User, error) {
	if name == "" || username == "" || email == "" || passwordHash == "" {
		return nil, ErrMissingRequiredField
	}

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Username:  username,
		Email:     email,
		Password:  passwordHash,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}, nil
}
