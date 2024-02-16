package pkg

import (
	"net/http"
)

var (
	ErrRecordNotFound = &RelayerError{
		Code:    http.StatusNotFound,
		Message: "Record Not Found",
	}

	ErrUnauthorized = &RelayerError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	ErrInternal = &RelayerError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Error",
	}
)

type RelayerError struct {
	Code    int32
	Message string
}

func (r *RelayerError) Error() string {
	return r.Message
}
