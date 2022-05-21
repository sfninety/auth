package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID
	PhoneNumber      string
	PasswordHash     string
	Verified         bool
	Onboarded        bool
	DeviceIdentifier string
	Updated          time.Time
	Created          time.Time
}
