package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

const gracefulThreshold = 15 * time.Second

// Handle handles application shutdown
func Handle(ctx context.Context, shutdownFunc func(ctx context.Context)) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt, syscall.SIGTERM:
		log.Info().Msg("[NOTICE] SIGINT or SIGTERM received")
		tCtx, cancel := context.WithTimeout(ctx, gracefulThreshold)

		defer cancel()

		shutdownFunc(tCtx)
	}
}
