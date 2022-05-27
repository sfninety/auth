package verification

import (
	"context"
	"time"
)

const (
	updateVerification = "UPDATE verifications SET otp = $2, expiry = $3 WHERE phone_number = $1;"
)

func (s *store) UpdateVerificationPair(ctx context.Context, v Verification) error {
	exp := time.Now().Add(time.Minute * 5)
	_, err := s.db.ExecContext(ctx, updateVerification, v.PhoneNumber, v.OTP, exp)
	return err
}
