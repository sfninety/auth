package user

import "context"

func (s *store) UpdateUser(ctx context.Context, user *User) error {
	return nil
}

func (s *store) VerifyUser(ctx context.Context, phoneNumber string) error {
	panic("not implemented!")
}
