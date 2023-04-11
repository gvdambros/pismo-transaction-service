package config

import (
	"fmt"
)

type (
	// Database db configuration
	Database struct {
		DatabaseDSN string
	}
)

func (c *cfg) parseDatabase() Database {
	return Database{
		DatabaseDSN: c.parseDatabaseInstance("rw"),
	}
}

func (c *cfg) parseDatabaseInstance(instance string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s connect_timeout=5",
		c.Get("db.host"),
		c.Get("db.user"),
		c.Get("db.password"),
		c.Get("db.name"),
		c.GetString("db.port"))
}
