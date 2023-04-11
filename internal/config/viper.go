package config

import (
	"fmt"
	"os"
	"strings"
	"time"
	"transaction-service/internal/config/files"

	"github.com/spf13/viper"
)

// cfg wrapper on viper for expanding/customizing default setup
type cfg struct{ *viper.Viper }

// load loads app's configuration from files and environment
func load() (*Configuration, error) {
	c := cfg{viper.New()}
	c.setFileLookup()
	c.setEnvLookup()
	c.setDefaults()

	if err := c.ReadInConfig(); err != nil {
		return nil, err
	}

	return c.parseConfig(), nil
}

// specifies file search path
func (c *cfg) setFileLookup() {
	c.AddConfigPath("./internal/config/files")
	c.AddConfigPath("./data/config/files")
	c.AddConfigPath(files.Path)
	c.SetConfigType("yaml")
	c.SetConfigName(c.getConfigFilename())
}

// configures properties read from environment
func (c *cfg) setEnvLookup() {
	c.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	c.AutomaticEnv()
}

// retrieves a configuration file name based on profile
func (c *cfg) getConfigFilename() string {
	profile := os.Getenv("GO_PROFILE")
	if profile != "" {
		return fmt.Sprintf("config-%s", strings.ToLower(profile))
	}

	return "config"
}

// retrieves a string key falling back to a default one if non-existent
func (c *cfg) getStringWithFallback(key, fallback string) string {
	if c.IsSet(key) {
		return c.GetString(key)
	}

	return c.GetString(fallback)
}

// retrieves an integer key falling back to a default one if non-existent
func (c *cfg) getIntWithFallback(key, fallback string) int {
	if c.IsSet(key) {
		return c.GetInt(key)
	}

	return c.GetInt(fallback)
}

// retrieves a boolean key falling back to a default one if non-existent
func (c *cfg) getBoolWithFallback(key, fallback string) bool {
	if c.IsSet(key) {
		return c.GetBool(key)
	}

	return c.GetBool(fallback)
}

// retrieves a numeric key in duration format (seconds) falling back to a default one if non-existent
func (c *cfg) getSecondsWithFallback(key, fallback string) time.Duration {
	if c.IsSet(key) {
		return c.getSeconds(key)
	}

	return c.getSeconds(fallback)
}

// retrieves a string slice key falling back to a default one if non-existent
func (c *cfg) getStringSliceWithFallback(key, fallback string) []string {
	if c.IsSet(key) {
		return c.GetStringSlice(key)
	}

	return c.GetStringSlice(fallback)
}

// retrieves a numeric key in duration format (seconds)
func (c *cfg) getSeconds(key string) time.Duration {
	return time.Duration(c.GetUint32(key)) * time.Second
}

// retrieves a numeric key in duration format (milliseconds)
func (c *cfg) getMilliseconds(key string) time.Duration {
	return time.Duration(c.GetUint32(key)) * time.Millisecond
}

func (c *cfg) setDefaults() {
	c.setApplicationDefaults()
	// c.setDatabaseDefaults()
	c.setDocsDefaults()
}

func (c *cfg) parseConfig() *Configuration {
	return &Configuration{
		Application:   c.parseApplication(),
		Database:      c.parseDatabase(),
		Documentation: c.parseDocs(),
	}
}
