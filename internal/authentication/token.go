package authentication

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sfninety/auth/ex/api"
	"github.com/sfninety/auth/internal/cryptography"
)

type AccessToken struct {
	EntropyBucket string `json:"eb"`
	jwt.RegisteredClaims
}

type AccessPair struct {
	Access  string
	Refresh string
}

func (ap *AccessPair) API() *api.AccessPairResponse {
	return &api.AccessPairResponse{
		AccessToken:  ap.Access,
		RefreshToken: ap.Refresh,
	}
}

func GenerateJwtPair(ctx context.Context, sk, encryptionKey, sub string) (*AccessPair, error) {
	iss := jwt.NewNumericDate(time.Now())

	at_eb := cryptography.GenerateEntropyBucket(10)
	at_exp := time.Now().Add(time.Minute * 5)

	at_jti := accessTokenPrefix + cryptography.GenerateEntropyBucket(6)

	at := AccessToken{
		EntropyBucket: at_eb,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        at_jti,
			ExpiresAt: jwt.NewNumericDate(at_exp),
			Audience:  jwt.ClaimStrings{"auth"},
			IssuedAt:  iss,
		},
	}

	rt_eb := cryptography.GenerateEntropyBucket(10)
	rt_exp := time.Now().Add(time.Hour * 24 * 7 * 31 * 12)

	rt_jti := refreshTokenPrefix + cryptography.GenerateEntropyBucket(6)

	rt := AccessToken{
		EntropyBucket: rt_eb,

		RegisteredClaims: jwt.RegisteredClaims{
			ID:        rt_jti,
			ExpiresAt: jwt.NewNumericDate(rt_exp),
			Audience:  jwt.ClaimStrings{"auth"},
			IssuedAt:  iss,
		},
	}

	signingKey := []byte(sk)

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, at)
	access_token_string, err := access_token.SignedString(signingKey)
	if err != nil {
		log.Printf("error signing access token: %v", err.Error())
		return nil, err
	}

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, rt)
	refresh_token_string, err := refresh_token.SignedString(signingKey)
	if err != nil {
		log.Printf("error signing refresh token: %v", err.Error())
		return nil, err
	}

	encrypted_at, err := cryptography.AESEncrypt(access_token_string, encryptionKey)
	if err != nil {
		log.Printf("failed to encrypt access token: %v", err)
		return nil, err
	}

	encrypted_rt, err := cryptography.AESEncrypt(refresh_token_string, encryptionKey)
	if err != nil {
		log.Printf("failed to encrypt refresh token: %v", err)
		return nil, err
	}

	// TODO: store the tokens, with the EB value, the SUB, and the EXP date

	return &AccessPair{
		Access:  encrypted_at,
		Refresh: encrypted_rt,
	}, nil
}
