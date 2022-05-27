package user

import (
	"context"
	"time"
)

const (
	createUser = "INSERT INTO users (phone_number, device_identifier, password_hash, created, updated) VALUES ($1, $2, $3, $4, $5);"
)

func (s *store) NewUser(ctx context.Context, phoneNumber, passwordHash, deviceIdentifier string) error {
	now := time.Now()
	_, err := s.db.ExecContext(ctx, createUser, phoneNumber, deviceIdentifier, passwordHash, now, now)
	return err
}
