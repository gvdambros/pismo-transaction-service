package echo

import (
	"strings"
	httperror "transaction-service/internal/pkg/http/server/echo/error"
	"transaction-service/internal/pkg/http/server/echo/validator"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const defaultHealthcheckTimeout = 10 // seconds
var ignoredPromPaths = []string{"/health", "/liveness", "/readiness", "/metrics", "/prometheus"}

func (s *Server) setup() {
	s.Echo = echo.New()
	s.Echo.HideBanner = false
	s.Echo.HidePort = false
	s.setMiddlewares()
}

func (s *Server) setMiddlewares() {
	s.Echo.HTTPErrorHandler = httperror.Handler
	s.Echo.Use(middleware.Recover())
	s.setCompression()
	s.setPrometheus()
	s.Echo.Validator = validator.New() // custom validator
	s.setLogger()
	s.setDebugLogging()
}

func (s *Server) setCompression() {
	s.Echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		// NOTE: gzip encoding breaks prometheus metrics gathering, so we skip it
		Skipper: func(c echo.Context) bool {
			requestURI := c.Request().RequestURI
			isPrometheus := requestURI == "/prometheus"
			isSwagger := strings.Contains(c.Request().URL.Path, "swagger")
			return isPrometheus || isSwagger
		},
	}))
}

func (s *Server) setDebugLogging() {
	// we want to log every incoming request when debugging (except for healthchecks and metrics)
	// NOTE: default mode for sandbox deployments
	s.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Debug().Msg("%+v")
			return nil
		},
		Skipper: func(c echo.Context) bool {
			requestURI := c.Request().RequestURI
			return requestURI == "/prometheus" || requestURI == "/liveness" || requestURI == "/readiness"
		},
	}))
}

func (s *Server) setPrometheus() {
	p := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
		basePath := c.Path()
		for _, i := range ignoredPromPaths {
			if strings.HasPrefix(basePath, i) {
				return true
			}
		}

		return false
	})
	p.Use(s.Echo)

	// metrics endpoint
	s.Echo.GET("/prometheus", echo.WrapHandler(promhttp.Handler()))
}

func (s *Server) setLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
