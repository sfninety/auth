package datastore

import (
	"database/sql"

	"github.com/sfninety/auth/internal/datastore/jti"
	"github.com/sfninety/auth/internal/datastore/token"
	"github.com/sfninety/auth/internal/datastore/user"
)

type store struct {
	Users  user.UserStore
	Tokens token.TokenStore
	Jtis   jti.JtiStore
}

type Config struct {
	ConnectionString string
}

var (
	cfg *Config

	Store *store
)

// Init establishes a database connection
func Init() {
	conn, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		panic(err)
	}

	Store = &store{
		Users:  user.Init(conn),
		Jtis:   jti.Init(conn),
		Tokens: token.Init(conn),
	}
}
