package server

import (
	"context"
)

// Server a http server
type Server interface {
	Start()
	Stop(context.Context)
	Register(...API)
	URL() string
}

// API a server API
type API interface {
	Routes() []Route
}

// Route an API route
type Route struct {
	Method  string
	Path    string
	Handler interface{}
}
