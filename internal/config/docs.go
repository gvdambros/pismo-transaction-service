package config

// Documentation openapi docs configuration
type Documentation struct {
	Enabled bool
}

// DefaultDocEnabled docs are enabled by default
const DefaultDocEnabled = true

func (c *cfg) parseDocs() Documentation {
	return Documentation{Enabled: c.GetBool("docs.enabled")}
}

func (c *cfg) setDocsDefaults() {
	c.SetDefault("docs.enabled", DefaultDocEnabled)
}
