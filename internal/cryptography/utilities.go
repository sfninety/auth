package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"io"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const (
	otpChars = "0123456789"
	allChars = otpChars + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bcryptHashLength = 14
)

func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = otpChars[rand.Intn(len(otpChars))]
	}
	return string(b)
}

func GenerateEntropyBucket(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = allChars[rand.Intn(len(allChars))]
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

func AESEncrypt(text, key string) (string, error) {
	c, err := aes.NewCipher([]byte(key)[:32])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(crypto_rand.Reader, nonce); err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(text), nil))

	return b64, nil
}

// TODO: decrypt token
