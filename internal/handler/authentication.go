// All the methods regarding AUTHENTICATING a user reside here.
package handler

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/sfninety/auth/ex/api"
	"github.com/sfninety/auth/internal/cryptography"
	"github.com/sfninety/auth/internal/datastore"
	"github.com/sfninety/auth/internal/httperrors"
	"github.com/sfninety/iris"
)

var (
	errInvalidPhoneNumber = errors.New("invalid phone number")
	errUserAlreadyExists  = errors.New("user already exists")
)

// This endpoint will ONLY ever be accessed by mobile devices.
// Register creates a new user and returns an access token and refresh token.
func Register(r iris.Request) iris.Response {
	req := &api.PhoneNumberRequest{}
	err := r.Decode(req)
	if err != nil || req.PhoneNumber == "" {
		return httperrors.BadRequest(r, "invalid request")
	}

	phoneNumber, err := parsePhoneNumber(req.PhoneNumber)
	if err != nil {
		return httperrors.BadRequest(r, err.Error())
	}

	ctx := r.Context

	_, err = datastore.Users.GetUser(ctx, phoneNumber)
	switch err {
	case sql.ErrNoRows:
		_, err = datastore.Users.NewUser(ctx, phoneNumber, "")
		if err != nil {
			log.Printf("failed to create new user: %v", err.Error())
			return httperrors.Internal(r, err.Error())
		}
	case nil:
		return httperrors.BadRequest(r, errUserAlreadyExists.Error())
	default:
		log.Printf("failed to query for existing user: %v", err.Error())
		return httperrors.Internal(r, err.Error())
	}

	otp := cryptography.GenerateOTP(6)
	_, err = datastore.Verifications.StoreVerificationPair(ctx, phoneNumber, otp)

	// TODO handle duplicates
	if err != nil {
		log.Printf("failed to create new otp: %v", err.Error())
		return httperrors.Internal(r, err.Error())
	}

	// TODO Send message through Kafka to notifications service to send a OTP text

	return r.ResponseWithCode(nil, 200)
}

func FinishRegistration(r iris.Request) iris.Response {
	req := api.RegistrationRequest{}
	err := r.Decode(&req)
	if err != nil || req.PhoneNumber == "" || req.Password == "" {
		return httperrors.BadRequest(r, "invalid request")
	}

	phoneNumber, err := parsePhoneNumber(req.PhoneNumber)
	if err != nil {
		return httperrors.BadRequest(r, err.Error())
	}

	ctx := r.Context

	user, err := datastore.Users.GetUser(ctx, phoneNumber)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		return httperrors.BadRequest(r, "no user associated with phone number")
	default:
		return httperrors.Internal(r, "failed to get user")
	}

	hashedPassword, err := cryptography.HashPassword(req.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err.Error())
		return httperrors.Internal(r, err.Error())
	}

	user.PasswordHash = hashedPassword
	user.Onboarded = true

	err = datastore.Users.UpdateUser(ctx, user)
	if err != nil {
		return httperrors.Internal(r, "failed to update user")
	}

	return r.ResponseWithCode(nil, 200)
}

func parsePhoneNumber(pn string) (string, error) {
	split := strings.Split(pn, "")

	if l := len(split); l > 14 || l < 13 {
		return "", errInvalidPhoneNumber
	}

	if strings.Join(split[:3], "") != "+44" {
		return "", errInvalidPhoneNumber
	}

	if split[3] == "0" {
		split = append(split[:3], split[4:]...)
	}

	return strings.Join(split, ""), nil
}
