package repository

import (
	"context"
	account "transaction-service/internal/domain/account"
	dbmodel "transaction-service/internal/domain/account/db/model"
)

func (r *Repository) SaveAccount(ctx context.Context, request *account.PostAccountRequest) (*account.Account, error) {
	account := r.convertDomainToDB(request)

	result := r.database.Connection.Create(account)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.convert(account), nil
}

func (r *Repository) GetAccount(ctx context.Context, request *account.GetAccountRequest) (*account.Account, error) {
	var account dbmodel.Account
	result := r.database.Connection.First(&account, request.Id)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.convert(&account), nil
}
