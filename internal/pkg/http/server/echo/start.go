package echo

import (
	"errors"
	"fmt"
	"net/http"
	"transaction-service/internal/config"

	"github.com/rs/zerolog/log"
)

// Start starts http server
func (s *Server) Start() {
	app := config.Get().Application
	name := app.Name
	port := app.Port

	log.Info().Msg(fmt.Sprintf("starting %s on %s...", name, port))

	go func() {
		if err := s.Echo.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msg("shutting down http server")
		}
	}()
}
