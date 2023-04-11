package config

// Application name and port for running application
type Application struct {
	Name string
	Port string
	Host string // for swagger
}

func (c *cfg) parseApplication() Application {
	return Application{
		Name: c.GetString("application.name"),
		Port: c.GetString("application.port"),
		Host: c.GetString("application.host"),
	}
}

func (c *cfg) setApplicationDefaults() {
	c.SetDefault("application.port", ":8080")
}
