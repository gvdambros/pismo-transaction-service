package echo

import (
	"fmt"
	"transaction-service/internal/pkg/http/server"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Register registers an echo api
func (s *Server) Register(collection ...server.API) {
	log.Info().Msg("registering routes:")

	for _, c := range collection {
		for _, route := range c.Routes() {
			log.Info().Msg(fmt.Sprintf("[%s] %s", route.Method, route.Path))
			m := []string{route.Method}

			switch h := route.Handler.(type) {
			case echo.HandlerFunc:
				s.Match(m, route.Path, h)
			case func(echo.Context) error:
				s.Match(m, route.Path, h)
			default:
				log.Fatal().Msg("couldn't register route")
			}
		}
	}
}
