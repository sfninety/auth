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
	GetUser(context.Context, string) (*User, error)
	NewUser(context.Context, string, string) (*User, error)
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, string) error
}

type store struct {
	db *sql.DB
}

func (s *store) GetUser(ctx context.Context, phoneNumber string) (*User, error) {
	return nil, nil
}

func (s *store) NewUser(ctx context.Context, phoneNumber string, passwordHash string) (*User, error) {
	return nil, nil
}

func (s *store) UpdateUser(ctx context.Context, user *User) error {
	return nil
}

func (s *store) DeleteUser(ctx context.Context, phoneNumber string) error {
	return nil
}
