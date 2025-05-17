package util

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/internal"
)

func HumaError(err error) error {
	var code int = http.StatusInternalServerError
	errMsg := "internal server error"

	var ierr *internal.Error
	if !errors.As(err, &ierr) {
		errMsg = "internal server error"
	} else {
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			code = http.StatusNotFound
			errMsg = ierr.Top().Error()

		case internal.ErrorCodeInvalidArgument:
			code = http.StatusBadRequest
			errMsg = ierr.Top().Error()

		case internal.ErrorCodeNotAuthorized:
			code = http.StatusUnauthorized
			errMsg = ierr.Top().Error()

		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			slog.Debug("error", "service error:", ierr.Error())
			code = http.StatusInternalServerError
			errMsg = "internal server error"
		}
	}

	return huma.NewError(code, errMsg)
}