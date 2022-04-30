package token

import "database/sql"

type TokenStore interface {
}

type store struct {
	db *sql.DB
}

func Init(db *sql.DB) TokenStore {
	return &store{
		db: db,
	}
}
