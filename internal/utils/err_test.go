package utils

import (
	"errors"
	"testing"

	app_errors "github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/stretchr/testify/assert"
)

func TestConvertHttpErrToGRPC(t *testing.T) {
	testcases := []struct {
		it       string
		appError *app_errors.AppError
		expected error
	}{
		{
			it:       "returns error with mapped code if the code is existed in the mapping",
			appError: app_errors.NotFoundError("something went wrong"),
			expected: errors.New("rpc error: code = NotFound desc = something went wrong"),
		},
		{
			it: "returns internal error if the code isn't existed in the mapping",
			appError: &app_errors.AppError{
				StatusCode: -1,
				Message:    "something went wrong",
			},
			expected: errors.New("rpc error: code = Internal desc = something went wrong"),
		},
	}

	for _, tC := range testcases {
		t.Run(tC.it, func(t *testing.T) {
			err := ConvertHttpErrToGRPC(tC.appError)

			assert.EqualError(t, err, tC.expected.Error())
		})
	}
}
