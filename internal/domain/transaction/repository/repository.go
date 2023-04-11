package repository

import (
	transaction "transaction-service/internal/domain/transaction"
	database "transaction-service/internal/pkg/database/postgres"
)

// Transactions repository
type Repository struct {
	database database.Database
}

// Transactions repository reference
func New() transaction.Repository {
	return &Repository{
		database: database.GetDatabase(),
	}
}
