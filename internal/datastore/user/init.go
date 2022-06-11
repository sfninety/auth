package user

import (
	"context"
	"database/sql"
)

func Init(db *sql.DB) UserStore {
	return &store{
		db: db,
	}
}

type UserStore interface {
	GetUser(ctx context.Context, phoneNumber string) (*User, error)
	NewUser(ctx context.Context, phoneNumber, passwordHash, deviceIdentifier string) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, phoneNumber string) error
	VerifyUser(ctx context.Context, phoneNumber string) error
	OnboardUser(ctx context.Context, devId, pwh, phoneNumber string) error
}

type store struct {
	db *sql.DB
}
