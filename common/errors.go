package common

import (
	"net/http"
)

//CustomError custom error inteerface
type CustomError interface {
	Message() string
	Status() int
	Causes() []interface{}
}

type customError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e customError) Message() string {
	return e.ErrMessage
}

func (e customError) Status() int {
	return e.ErrStatus
}

func (e customError) Causes() []interface{} {
	return e.ErrCauses
}

//NewBadRequestError new bad request (usually bad data)
func NewBadRequestError(message string) CustomError {
	return customError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError not found error
func NewNotFoundError(message string) CustomError {
	return customError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

//NewInternalServerError internal server error
func NewInternalServerError(message string, err error) CustomError {
	result := customError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}

	return result
}

//NewUnauthorizedError unauthorized
func NewUnauthorizedError(message string) CustomError {
	return customError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}
