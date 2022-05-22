package datastore

import (
	"database/sql"

	"github.com/sfninety/auth/internal/datastore/jti"
	"github.com/sfninety/auth/internal/datastore/token"
	"github.com/sfninety/auth/internal/datastore/user"
	"github.com/sfninety/auth/internal/datastore/verification"
)

var (
	Users         user.UserStore
	Tokens        token.TokenStore
	Jtis          jti.JtiStore
	Verifications verification.VerificationStore
)

type Config struct {
	ConnectionString string `yaml:"connection_string"`
}

// Init establishes a database connection
func Init(cfg *Config) {

	conn, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		panic(err)
	}

	Users = user.Init(conn)
	Jtis = jti.Init(conn)
	Tokens = token.Init(conn)
	Verifications = verification.Init(conn)
}
