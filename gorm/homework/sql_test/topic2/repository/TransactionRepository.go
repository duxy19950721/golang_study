package repository

import (
	"golang_study/gorm/homework/sql_test/topic2/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

// 保存交易记录
func (tr *TransactionRepository) RecordTransaction(fromAccountId string, toAccountId string, amount int) error {
	transcation := model.Transaction{FromAccountId: fromAccountId, ToAccountId: toAccountId, Amount: amount}
	result := tr.DB.Create(&transcation)
	return result.Error
}
