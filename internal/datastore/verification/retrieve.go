package verification

import "context"

const (
	retrieveVerificationPair = "SELECT * FROM verifications WHERE phone_number = $1;"
)

func (s *store) RetrieveVerificationPair(ctx context.Context, phoneNumber string) (*Verification, error) {
	verification := &Verification{}
	err := s.db.QueryRowContext(ctx, retrieveVerificationPair, phoneNumber).Scan(&verification.Id, &verification.PhoneNumber, &verification.OTP, &verification.Exp)
	return verification, err
}
