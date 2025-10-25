package repository

import (
	"golang_study/gorm/homework/sql_test/topic2/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

// 根据id查询账户信息
func (ar *AccountRepository) GetAccountById(id string) *model.Account {
	account := model.Account{}
	ar.DB.First(&account, "id = ?", id)
	return &account
}

// 更新账户余额
func (ar *AccountRepository) UpdateAccount(id string, amount int) error {
	account := ar.GetAccountById(id)
	result := ar.DB.Model(&account).Select("balance").Update("balance", account.Balance-int(amount))
	return result.Error
}
