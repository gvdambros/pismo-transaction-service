package echo

import (
	"transaction-service/internal/pkg/http/server"

	"github.com/labstack/echo/v4"
)

// Server an echo http server
type Server struct {
	*echo.Echo
	url string
}

// New retrieves a reference to an echo http server
func New() server.Server {
	s := &Server{}
	s.setup()

	return s
}
