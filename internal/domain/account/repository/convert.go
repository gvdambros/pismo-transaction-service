package repository

import (
	account "transaction-service/internal/domain/account"
	dbmodel "transaction-service/internal/domain/account/db/model"
)

func (r *Repository) convertDomainToDB(request *account.PostAccountRequest) *dbmodel.Account {
	return &dbmodel.Account{
		DocumentNumber: request.DocumentNumber,
	}
}

func (r *Repository) convert(res *dbmodel.Account) *account.Account {
	return &account.Account{
		Id:             res.Id,
		DocumentNumber: res.DocumentNumber,
	}
}
