package main

import (
	"fmt"

	"github.com/jimmytong/golang-task/task3/assignment2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//	assignment2
	db.AutoMigrate(&assignment2.Account{}, &assignment2.Transaction{})
	A := assignment2.Account{Balance: 1000.00}
	B := assignment2.Account{Balance: 200.00}
	db.Create(&A)
	db.Create(&B)
	fmt.Printf("账户A初始余额: %.2f\n", A.Balance)
	fmt.Printf("账户B初始余额: %.2f\n", B.Balance)

	err = assignment2.TransferMoney(db, A.ID, B.ID, 1100.00)
	if err != nil {
		panic(err)
	}
	fmt.Println("转账成功")

	// assignment1.Run(db)
	// assignment1.CreateStudents(db)
	// assignment1.QueryStudents(db)
	// assignment1.DeleteStudents(db)
	// assignment1.Field(db)
}
