package it

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"time"
	account "transaction-service/internal/domain/account/api"
	transaction "transaction-service/internal/domain/transaction/api"
	"transaction-service/internal/pkg/http/server/echo"

	"github.com/rs/zerolog/log"
)

const (
	serviceWait    = 2 * time.Second
	serviceTimeout = 30 * time.Second
)

func setup() {
	// API
	localServer = echo.New()
	localServer.Register(account.New(), transaction.New())
	localServer.Start()
}

func waitSingleService(service, port string) {
	ctx := context.Background()
	log.Info().Msgf("[%s] waiting...", service)

	tCtx, cancel := context.WithTimeout(ctx, serviceTimeout)
	defer cancel()

	for {
		if errors.Is(tCtx.Err(), context.Canceled) {
			log.Info().Msgf("[%s] timed out", service)
		}

		connString := fmt.Sprintf("localhost:%s", port)

		conn, err := net.Dial("tcp", connString)
		if err != nil {
			time.Sleep(serviceWait)
			continue
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "started\n" {
			log.Info().Msgf("[%s] ready!", service)
			return
		}

		log.Info().Msgf("[%s] retrying...", service)
		time.Sleep(serviceWait)
	}
}
