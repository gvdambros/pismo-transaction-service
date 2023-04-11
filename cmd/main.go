package main

import (
	"context"
	"transaction-service/internal/config"
	account "transaction-service/internal/domain/account/api"
	transaction "transaction-service/internal/domain/transaction/api"
	"transaction-service/internal/pkg/http/server/echo"
	"transaction-service/internal/pkg/shutdown"
	"transaction-service/internal/pkg/swagger"
)

// @title Transaction Service
// @version 1.0
// @description API supporting accounts and its transactions
func main() {
	globalCtx := context.Background()

	// http server
	sv := echo.New()
	sv.Register(
		account.New(),
		transaction.New(),
	)

	// swagger
	if config.Get().Documentation.Enabled {
		sv.Register(swagger.New())
	}

	sv.Start()

	// gracefully handles shutdowns
	shutdown.Handle(globalCtx, func(ctx context.Context) {
		sv.Stop(ctx)
	})
}
