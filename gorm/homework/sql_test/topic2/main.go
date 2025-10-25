package main

import (
	"errors"
	"fmt"
	"golang_study/gorm/homework/sql_test/config"
	"golang_study/gorm/homework/sql_test/topic2/model"
	"golang_study/gorm/homework/sql_test/topic2/repository"

	"gorm.io/gorm"
)

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）
// 和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

func main() {
	db := config.InitDB()
	// accountRepository := repository.AccountRepository{DB: db}
	// transactionRepository := repository.TransactionRepository{DB: db}

	// 初始化账户信息
	db.AutoMigrate(&model.Account{})
	// 初始化交易记录表结构
	db.AutoMigrate(&model.Transaction{})

	// 设置转账金额
	amount := 100

	error := db.Transaction(func(tx *gorm.DB) error {
		if amount <= 0 {
			return errors.New("amount must be greater than 0")
		}

		accountRepository := repository.AccountRepository{DB: tx}
		transactionRepository := repository.TransactionRepository{DB: tx}

		// 先查询A的账户余额信息
		fromAccount := accountRepository.GetAccountById("A")
		if fromAccount.Balance < amount {
			return errors.New("余额不足,无法转账")
		}
		// 开始转账，A-100，B+100
		accountRepository.UpdateAccount("A", amount)
		accountRepository.UpdateAccount("B", -amount)

		// 保存交易记录
		error := transactionRepository.RecordTransaction("A", "B", amount)
		if error != nil {
			return error
		}

		fmt.Println("转账成功")

		return nil
	})
	if error != nil {
		fmt.Println("出现错误:", error)
	}
}
