package verification

import "time"

type Verification struct {
	PhoneNumber string
	OTP         string
	Exp         time.Time
}
