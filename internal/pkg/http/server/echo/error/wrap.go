package error

import (
	"errors"
	"net/http"
	httperror "transaction-service/internal/pkg/http/error"

	"github.com/labstack/echo/v4"
)

// wraps an unknown or unmapped error
func wrap(err error) httperror.Wrapper {
	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		switch echoErr.Code {
		case http.StatusNotFound:
			return httperror.Build(httperror.NotFound, echoErr)
		default:
			return httperror.Build(httperror.BadRequest, echoErr)
		}
	}

	return httperror.Build(httperror.ServerError, err)
}
