package database

import (
	"context"
	"sync"
	"transaction-service/internal/config"

	"github.com/joomcode/errorx"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	runOnceConnect sync.Once
	database       *Database
)

func GetDatabase() Database {
	runOnceConnect.Do(func() {
		database = &Database{
			Connection: nil,
			err:        nil,
		}
	})

	if database.Connection == nil {
		database.Connect()
	}

	return *database
}

type Database struct {
	Connection *gorm.DB
	err        error
}

func (c *Database) Connect() {
	dsn := config.Get().Database.DatabaseDSN
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Failed to connect to database")
	}
	c.Connection = connection
	c.err = err
}

func (c Database) Ping(ctx context.Context) *errorx.Error {
	results, err := c.Connection.DB()
	if err != nil {
		return errorx.Decorate(err, "Unable to get connection")
	}

	if pingError := results.Ping(); pingError != nil {
		return errorx.Decorate(err, "Unable to ping")
	}
	return nil
}
