package user

import (
	"context"
	"time"
)

const (
	verifyUser  = "UPDATE users SET verified = true, updated = $2 WHERE phone_number = $1;"
	onboardUser = "UPDATE users SET device_identifier = $2, password_hash = $3, updated = $4 WHERE phone_number = $1;"
)

func (s *store) UpdateUser(ctx context.Context, user *User) error {
	return nil
}

func (s *store) OnboardUser(ctx context.Context, devId, pwh, phoneNumber string) error {
	_, err := s.db.ExecContext(ctx, onboardUser, phoneNumber, devId, pwh, time.Now())
	return err
}

func (s *store) VerifyUser(ctx context.Context, phoneNumber string) error {
	_, err := s.db.ExecContext(ctx, verifyUser, phoneNumber, time.Now())
	return err
}
