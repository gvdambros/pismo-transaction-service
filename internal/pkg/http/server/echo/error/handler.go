package error

import (
	"errors"
	"net/http"
	httperror "transaction-service/internal/pkg/http/error"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Handler custom echo error handler
func Handler(err error, c echo.Context) {
	var custom *httperror.Error
	if !errors.As(err, &custom) {
		custom = wrap(err).Custom()
	}

	status := custom.Status
	path := c.Request().RequestURI

	if status < http.StatusInternalServerError {
		// client errors
		log.Warn().Msgf("[%s] user error: %+v", path, err)
	} else {
		// server errors
		log.Error().Msgf("[%s] server error: %s", path, err.Error())
	}

	if !c.Response().Committed {
		_ = c.JSON(status, custom)
	}
}
