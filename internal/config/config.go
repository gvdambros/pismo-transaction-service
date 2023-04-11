package config

import (
	"sync"

	"github.com/rs/zerolog/log"
)

// Configuration defines application's configurable aspects
type Configuration struct {
	Application
	Database
	Documentation
}

var (
	configOnce sync.Once
	config     *Configuration
)

// Get retrieves global configuration
func Get() *Configuration {
	var err error
	configOnce.Do(func() { config, err = load() })

	if err != nil {
		log.Fatal().Msg("failed to load app's configuration")
	}

	return config
}
