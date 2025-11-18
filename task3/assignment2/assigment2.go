package assignment2

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	ID      uint
	Balance float64 `gorm:"not null"`
}

type Transaction struct {
	ID            uint
	FromAccountId uint    `gorm:"not null"`
	ToAccountId   uint    `gorm:"not null"`
	Account       float64 `gorm:"not null"`
}

func TransferMoney(db *gorm.DB, fromAccountId, toAccountId uint, amount float64) error {
	fmt.Println("开始转账交易....")
	return db.Transaction(func(tx *gorm.DB) error {
		// var fromAccount, toAccount Account
		fromAccount := Account{
			ID: fromAccountId,
		}
		toAccount := Account{
			ID: toAccountId,
		}

		// 检查账户 A 的余额是否足够
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Debug().First(&fromAccount).Error; err != nil {
			fmt.Println("查询账户A失败:", err)
			return err
		}
		if fromAccount.Balance < amount {

			return fmt.Errorf("账户 %d 余额不足", fromAccountId)
		}

		// 扣除账户 A 的余额
		fromAccount.Balance -= amount
		if err := tx.Save(&fromAccount).Error; err != nil {
			fmt.Println("更新账户A失败:", err)
			return err
		}

		// 增加账户 B 的余额
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&toAccount).Error; err != nil {
			fmt.Println("查询账户B失败:", err)
			return err
		}
		toAccount.Balance += amount
		if err := tx.Save(&toAccount).Error; err != nil {
			fmt.Println("更新账户B失败:", err)
			return err
		}

		// 记录转账信息
		transaction := Transaction{
			FromAccountId: fromAccountId,
			ToAccountId:   toAccountId,
			Account:       amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			fmt.Println("创建转账记录失败:", err)
			return err
		}

		return nil
	})
}
