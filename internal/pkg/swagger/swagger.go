package swagger

import (
	"net/http"
	"transaction-service/internal/config"
	"transaction-service/internal/gen/doc"
	"transaction-service/internal/pkg/http/server"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// Swagger a swagger documentation handler
type Swagger struct {
}

func (s *Swagger) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodGet, Path: "/swagger/*", Handler: echoSwagger.WrapHandler},
	}
}

// New a new swagger documentation handler reference
func New() *Swagger { //nolint:wsl
	doc.SwaggerInfo.Host = config.Get().Application.Host

	return &Swagger{}
}
