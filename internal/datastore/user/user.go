package user

import "time"

type User struct {
	Id           int32
	PhoneNumber  string
	PasswordHash string
	Updated      time.Time
	Created      time.Time
}
