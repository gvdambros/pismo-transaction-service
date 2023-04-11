package echo

import (
	"context"

	"github.com/rs/zerolog/log"
)

// Stop gracefully stops http server
func (s *Server) Stop(ctx context.Context) {
	if err := s.Echo.Shutdown(ctx); err != nil {
		log.Panic().Msg("unresponsive: forcefully shutting down server")
	}
}
