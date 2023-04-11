package transaction

import (
	"context"
	"time"
)

type (
	Transaction struct {
		Id              int       `json:"id"`
		AccountId       int       `json:"account_id"`
		OperationTypeId int       `json:"operation_type_id"`
		Amount          float64   `json:"amount"`
		EventDate       time.Time `json:"event_date"`
	}

	PostTransactionRequest struct {
		AccountId       int     `json:"account_id" validate:"required,gt=0"`
		OperationTypeId int     `json:"operation_type_id" validate:"required,gt=0"`
		Amount          float64 `json:"amount" validate:"required,ne=0"`
	}

	// Handler a service for handling transactions
	Handler interface {
		Transaction(ctx context.Context, request *PostTransactionRequest) (*Transaction, error)
	}

	// Repository a service for persisting a transaction locally
	Repository interface {
		SaveTransaction(context.Context, *PostTransactionRequest) (*Transaction, error)
	}
)
