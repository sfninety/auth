package user

import "context"

func (s *store) NewUser(ctx context.Context, phoneNumber, passwordHash, deviceIdentifier string) (*User, error) {
	return nil, nil
}
