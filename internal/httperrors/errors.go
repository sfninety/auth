package httperrors

import (
	"net/http"

	"github.com/sfninety/iris"
)

type errorResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BadRequest(r iris.Request, body string) iris.Response {
	return r.ResponseWithCode(&errorResponseBody{
		Code:    http.StatusBadRequest,
		Message: body,
	}, http.StatusBadRequest)
}

func Internal(r iris.Request, body string) iris.Response {
	return r.ResponseWithCode(&errorResponseBody{
		Code:    http.StatusInternalServerError,
		Message: body,
	}, http.StatusInternalServerError)
}

func Unauthorized(r iris.Request, body string) iris.Response {
	return r.ResponseWithCode(&errorResponseBody{
		Code:    http.StatusUnauthorized,
		Message: body,
	}, http.StatusUnauthorized)
}
