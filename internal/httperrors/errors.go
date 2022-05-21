package httperrors

import (
	"net/http"

	"github.com/sfninety/iris"
)

type errorResponseBody struct {
	code    int
	message string
}

func BadRequest(r iris.Request, body string) iris.Response {
	return r.ResponseWithCode(&errorResponseBody{
		code:    http.StatusBadRequest,
		message: body,
	}, http.StatusBadRequest)
}

func Internal(r iris.Request, body string) iris.Response {
	return r.ResponseWithCode(&errorResponseBody{
		code:    http.StatusInternalServerError,
		message: body,
	}, http.StatusInternalServerError)
}

func Unauthorized(r iris.Request, body string) iris.Response {
	return r.ResponseWithCode(&errorResponseBody{
		code:    http.StatusUnauthorized,
		message: body,
	}, http.StatusUnauthorized)
}
