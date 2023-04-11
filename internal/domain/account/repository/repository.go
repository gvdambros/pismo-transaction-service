package repository

import (
	account "transaction-service/internal/domain/account"
	dbmodel "transaction-service/internal/domain/account/db/model"
	database "transaction-service/internal/pkg/database/postgres"
)

// Account repository
type Repository struct {
	database database.Database
}

type fetchOne struct {
	dbmodel.Account
}

// Account repository reference
func New() account.Repository {
	return &Repository{
		database: database.GetDatabase(),
	}
}
