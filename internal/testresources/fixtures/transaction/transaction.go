package transaction

import (
	"time"
	"transaction-service/internal/domain/transaction"
)

// Transaction fixture
func Transaction() *transaction.Transaction {

	return &transaction.Transaction{
		Id:              1,
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          1,
		EventDate:       time.Now(),
	}
}
