package handler

import (
	"database/sql"
	"log"

	"github.com/sfninety/auth/ex/api"
	"github.com/sfninety/auth/internal/cryptography"
	"github.com/sfninety/auth/internal/datastore"
	"github.com/sfninety/auth/internal/datastore/verification"
	"github.com/sfninety/auth/internal/httperrors"
	"github.com/sfninety/iris"
)

func RequestNewOTP(r iris.Request) iris.Response {
	req := &api.PhoneNumberRequest{}
	err := r.Decode(req)
	if err != nil {
		return r.ResponseWithCode("invalid request", 400)
	}

	phoneNumber, err := parsePhoneNumber(req.PhoneNumber)
	if err != nil || req.PhoneNumber == "" {
		return httperrors.BadRequest(r, "invalid request")
	}

	ctx := r.Context

	otp := cryptography.GenerateOTP(6)

	_, err = datastore.Verifications.RetrieveVerificationPair(ctx, phoneNumber)
	switch err {
	case sql.ErrNoRows:
		_, err = datastore.Verifications.StoreVerificationPair(ctx, phoneNumber, otp)
		if err != nil {
			log.Printf("failed to create new otp: %v", err.Error())
			return httperrors.Internal(r, "failed to create new otp")
		}
	case nil:
		err = datastore.Verifications.UpdateVerificationPair(ctx, verification.Verification{
			PhoneNumber: phoneNumber,
			OTP:         otp,
		})
		if err != nil {
			log.Printf("failed to create new otp: %v", err.Error())
			return httperrors.Internal(r, "failed to create new otp")
		}

	default:
		log.Printf("failed to retrieve verification pair: %v", err.Error())
		return httperrors.Internal(r, err.Error())
	}

	// TODO send message to notification server to send OTP text

	return r.ResponseWithCode(nil, 200)
}

func VerifyOTP(r iris.Request) iris.Response {
	req := &api.VerifyOTPRequest{}
	err := r.Decode(req)
	if err != nil || req.PhoneNumber == "" || req.Otp == "" {
		return httperrors.BadRequest(r, "invalid request")
	}

	phoneNumber, err := parsePhoneNumber(req.PhoneNumber)
	if err != nil {
		return httperrors.BadRequest(r, err.Error())
	}

	ctx := r.Context

	v, err := datastore.Verifications.RetrieveVerificationPair(ctx, phoneNumber)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		return httperrors.BadRequest(r, "no verification associated with phone number")
	default:
		log.Printf("failed to retrieve verification pair: %v", err)
		return httperrors.Internal(r, "failed to retrieve verification pair")
	}

	if v.OTP != req.Otp {
		return httperrors.Unauthorized(r, "otp does not match")
	}

	err = datastore.Users.VerifyUser(ctx, phoneNumber)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		return httperrors.BadRequest(r, "no user associated with phone number")
	default:
		log.Printf("failed to retrieve user: %v", err)
		return httperrors.Internal(r, "failed to verify user")
	}

	return r.ResponseWithCode(nil, 200)
}
