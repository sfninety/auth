package verification

import (
	"context"
	"time"
)

const (
	createVerification = "INSERT INTO verifications (phone_number, otp, expiry) VALUES ($1, $2, $3);"
)

func (s *store) StoreVerificationPair(ctx context.Context, phoneNumber string, otp string) error {
	exp := time.Now().Add(time.Minute * 5)
	_, err := s.db.ExecContext(ctx, createVerification, phoneNumber, otp, exp)
	return err
}
