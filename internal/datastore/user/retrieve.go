package user

import "context"

const (
	allFields    = "id, phone_number, device_identifier, verified, onboarded, created, updated"
	retrieveUser = "SELECT " + allFields + " FROM users WHERE phone_number = $1;"
)

func (s *store) GetUser(ctx context.Context, phoneNumber string) (*User, error) {
	user := &User{}
	err := s.db.QueryRowContext(ctx, retrieveUser, phoneNumber).Scan(&user.Id, &user.PhoneNumber, &user.DeviceIdentifier, &user.Verified, &user.Onboarded, &user.Created, &user.Updated)
	return user, err
}
