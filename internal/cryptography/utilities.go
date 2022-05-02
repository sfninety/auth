package cryptography

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const (
	otpChars = "0123456789"

	bcryptHashLength = 14
)

func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(b)
}

func HashPassword(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcryptHashLength)
	return string(bytes), err
}

func VerifyPassword(plaintext, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext)) == nil
}
