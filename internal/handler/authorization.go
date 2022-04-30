// All the methods regarding AUTHENTICATING a user reside here.
package handler

import "github.com/sfninety/iris"

type registerRequest struct {
	PhoneNumber       string
	PlaintextPassword string
	DeviceIdentifier  string
}

// This endpoint will ONLY ever be accessed by mobile devices.
// Register creates a new user and returns an access token and refresh token.
func Register(r iris.Request) iris.Response {
	req := &registerRequest{}
	err := r.Decode(req)
	if err != nil {
		return r.ResponseWithCode("invalid request", 400)
	}

	return r.ResponseWithCode(nil, 200)
}
