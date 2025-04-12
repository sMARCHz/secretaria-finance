package utils

import (
	"net/http"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var statusCodeMap = map[int]codes.Code{
	http.StatusBadRequest:          codes.InvalidArgument,
	http.StatusNotFound:            codes.NotFound,
	http.StatusUnprocessableEntity: codes.FailedPrecondition,
	http.StatusInternalServerError: codes.Internal,
}

func ConvertHttpErrToGRPC(appError *errors.AppError) error {
	statusCode, exist := statusCodeMap[appError.StatusCode]
	if !exist {
		statusCode = codes.Internal
	}
	return status.Error(statusCode, appError.Message)
}
