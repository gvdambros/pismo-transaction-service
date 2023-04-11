package api

import (
	"net/http"
	account "transaction-service/internal/domain/account"
	"transaction-service/internal/domain/account/handler"
	"transaction-service/internal/pkg/http/server"
)

// Account API
type API struct {
	handler account.Handler
}

// Account routes
func (a *API) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodPost, Path: "/accounts", Handler: a.Post},
		{Method: http.MethodGet, Path: "/accounts/:id", Handler: a.Get},
	}
}

// Account API reference
func New() server.API {
	return &API{handler: handler.NewAccount()}
}
