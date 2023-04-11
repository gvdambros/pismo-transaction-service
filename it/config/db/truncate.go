package db

import (
	"context"
	"fmt"
	database "transaction-service/internal/pkg/database/postgres"
)

// Truncate truncates tables in test db
func Truncate(ctx context.Context, tables ...string) error {
	for _, table := range tables {
		stmt := fmt.Sprintf("TRUNCATE table %s CASCADE;", table)
		if result := database.GetDatabase().Connection.Exec(stmt); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
