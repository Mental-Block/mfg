package util

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/server/pkg/utils"
)

func HumaError(err error) error {
	code := http.StatusInternalServerError
	errMsg := "internal server error"

	var ierr *utils.Error
	if errors.As(err, &ierr) {
		switch ierr.Code() {
		case utils.ErrorCodeNotFound:
			code = http.StatusNotFound
			errMsg = ierr.Top().Error()

		case utils.ErrorCodeInvalidArgument:
			code = http.StatusBadRequest
			errMsg = ierr.Top().Error()

		case utils.ErrorCodeNotAuthorized:
			code = http.StatusUnauthorized
			errMsg = ierr.Top().Error()

		case utils.ErrorCodeUnknown:
			fallthrough
		default:
			slog.Debug("error", "service error:", ierr.Error())
			code = http.StatusInternalServerError
			errMsg = "internal server error"
		}
	}

	return huma.NewError(code, errMsg)
}