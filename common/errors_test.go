package common

import (
	"errors"
	"testing"
)

func Test_NewBadRequestError(t *testing.T) {

	err := NewBadRequestError("message")

	if err.Status() != 400 {
		t.Error("Wrong status, expected 400, got ", err.Status())
	}

}

func Test_NewNotFoundError(t *testing.T) {

	err := NewNotFoundError("message")

	if err.Status() != 404 {
		t.Error("Wrong status, expected 404, got ", err.Status())
	}

}

func Test_NewInternalServerError(t *testing.T) {

	err := NewInternalServerError("message", errors.New(""))

	if err.Status() != 500 {
		t.Error("Wrong status, expected 500, got ", err.Status())
	}

}

func Test_NewUnauthorizedError(t *testing.T) {

	err := NewUnauthorizedError("message")

	if err.Status() != 401 {
		t.Error("Wrong status, expected 401, got ", err.Status())
	}

}
