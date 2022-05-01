package cryptography

import "math/rand"

const (
	otpChars = "0123456789"
)

func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(b)
}
