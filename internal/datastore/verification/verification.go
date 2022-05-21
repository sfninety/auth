package verification

import (
	"time"

	"github.com/google/uuid"
)

type Verification struct {
	Id          uuid.UUID
	PhoneNumber string
	OTP         string
	Exp         time.Time
}
