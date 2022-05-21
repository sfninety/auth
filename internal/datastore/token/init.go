package token

import (
	"context"
	"database/sql"
)

type TokenStore interface {
	CreateToken(ctx context.Context, eb, sub string, exp int32) error
	RetrieveToken(ctx context.Context, eb string) (*Token, error)
	DeleteToken(ctx context.Context, eb string) error
}

type store struct {
	db *sql.DB
}

func Init(db *sql.DB) TokenStore {
	return &store{
		db: db,
	}
}
