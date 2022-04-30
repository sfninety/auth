package jti

import "database/sql"

type JtiStore interface {
}

type store struct {
	db *sql.DB
}

func Init(db *sql.DB) JtiStore {
	return &store{
		db: db,
	}
}
