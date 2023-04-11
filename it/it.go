package it

import (
	"os"
	"sync"
	"transaction-service/internal/pkg/http/server"

	"github.com/onsi/gomega/ghttp"
)

var (
	localServer server.Server
	mockServer  *ghttp.Server
	initOnce    sync.Once
)

// Init ensures test setup is up
func Init() {
	os.Setenv("GO_PROFILE", "test")
	initOnce.Do(func() { setup() })
}

// LocalServer global reference of local server
func LocalServer() server.Server {
	return localServer
}
