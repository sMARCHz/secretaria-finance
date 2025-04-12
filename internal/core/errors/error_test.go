package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errorMsg = "something went wrong"

func TestBadRequestError(t *testing.T) {
	err := BadRequestError(errorMsg)

	expected := &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    errorMsg,
	}
	assert.Equal(t, expected, err)
}

func TestNotFoundError(t *testing.T) {
	err := NotFoundError(errorMsg)

	expected := &AppError{
		StatusCode: http.StatusNotFound,
		Message:    errorMsg,
	}
	assert.Equal(t, expected, err)
}

func TestUnprocessableEntityServerError(t *testing.T) {
	err := UnprocessableEntityServerError(errorMsg)

	expected := &AppError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errorMsg,
	}
	assert.Equal(t, expected, err)
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError(errorMsg)

	expected := &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    errorMsg,
	}
	assert.Equal(t, expected, err)
}
