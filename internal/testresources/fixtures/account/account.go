package account

import (
	"transaction-service/internal/domain/account"
)

// Account fixture
func Account() *account.Account {

	return &account.Account{
		Id:             1,
		DocumentNumber: "00427255007",
	}
}
