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
