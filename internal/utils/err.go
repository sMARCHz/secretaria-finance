package utils

import (
	"net/http"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
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
		logger.Errorf("appError status code['%v'] isn't in the map", appError.StatusCode)
		statusCode = codes.Internal
	}
	return status.Error(statusCode, appError.Message)
}
