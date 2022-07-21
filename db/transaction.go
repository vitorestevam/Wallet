package db

import (
	"gorm.io/gorm"
)

type TransactionKind int

const (
	Paying    TransactionKind = 0
	Receiving TransactionKind = 1
)

type TransactionBase struct {
	Name        string
	Description string
	Value       float64
	Kind        TransactionKind
}

type Transaction struct {
	gorm.Model
	TransactionBase
}

func CreateTransaction(database *gorm.DB, base TransactionBase) {
	database.Create(&Transaction{TransactionBase: base})
}

func GetTransactions(database *gorm.DB) []Transaction {
	var trans []Transaction
	database.Find(&trans)
	return trans
}

func DeleteTransaction(database *gorm.DB, id int) {
	database.Delete(&Transaction{}, id)
}
