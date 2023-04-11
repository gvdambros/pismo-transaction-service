package api

import (
	"net/http"
	transaction "transaction-service/internal/domain/transaction"
	"transaction-service/internal/domain/transaction/handler"
	"transaction-service/internal/pkg/http/server"
)

// Transaction API
type API struct {
	handler transaction.Handler
}

// Transaction routes
func (a *API) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodPost, Path: "/transactions", Handler: a.New},
	}
}

// Transaction API reference
func New() server.API {
	return &API{handler: handler.NewTransaction()}
}
