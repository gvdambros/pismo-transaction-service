package model

import "time"

type Transaction struct {
	Id              int `sql:"primary_key"`
	AccountId       int
	OperationTypeId int
	Amount          float64
	EventDate       time.Time
}
