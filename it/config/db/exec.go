package db

import (
	"context"
	database "transaction-service/internal/pkg/database/postgres"
)

// Exec execs a statement in test db
func Exec(ctx context.Context, statement *string) error {
	result := database.GetDatabase().Connection.Exec(*statement)

	return result.Error
}
