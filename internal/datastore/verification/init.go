package verification

import (
	"context"
	"database/sql"
)

type VerificationStore interface {
	StoreVerificationPair(ctx context.Context, phoneNumber string, otp string) (*Verification, error)
	RetrieveVerificationPair(ctx context.Context, phoneNumber string) (*Verification, error)
	DeleteVerificationPair(ctx context.Context, phoneNumber string) error
	UpdateVerificationPair(ctx context.Context, v Verification) error
}

type store struct {
	db *sql.DB
}

func Init(db *sql.DB) VerificationStore {
	return &store{
		db: db,
	}
}
