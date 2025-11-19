package main

import (
	"fmt"

	"github.com/jimmytong/golang-task/task3/assignment4"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	//assignment4任务四
	db := assignment4.InitDB()
	var count int64
	var id uint = 4
	db.Debug().Model(&assignment4.Post{}).Where("user_id", id).Count(&count)

	//创建一条用户数据和关联的文章数据，触发钩子函数
	newUser := assignment4.User{
		Name:  "测试用户",
		Email: "123@123.com",
		Posts: []assignment4.Post{
			{Title: "测试文章1", Content: "这是测试文章1的内容"},
			{Title: "测试文章2", Content: "这是测试文章2的内容"},
		},
	}
	db.Debug().Create(&newUser)

	// 查询用户的文章数量，验证钩子函数是否生效	
	user := assignment4.User{}
	user.ID = id
	db.Debug().First(&user)
	fmt.Println(user.PostCount)
	fmt.Printf("用户 %d 的文章数量: %d\n", id, count)
	// assignment4.MockData()
	

	// work.GetUserPostsAndComments(db, 1)
	// maxPost, err := work.GetMostCommentedPost(db)
	// if err != nil {
	// 	panic(err)
	// }
	// println("评论数量最多的文章标题:", maxPost.Title)

	//	assignment2
	// db.AutoMigrate(&assignment2.Account{}, &assignment2.Transaction{})
	// A := assignment2.Account{Balance: 1000.00}
	// B := assignment2.Account{Balance: 200.00}
	// db.Create(&A)
	// db.Create(&B)
	// fmt.Printf("账户A初始余额: %.2f\n", A.Balance)
	// fmt.Printf("账户B初始余额: %.2f\n", B.Balance)

	// err = assignment2.TransferMoney(db, A.ID, B.ID, 1100.00)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("转账成功")

	// assignment1.Run(db)
	// assignment1.CreateStudents(db)
	// assignment1.QueryStudents(db)
	// assignment1.DeleteStudents(db)
	// assignment1.Field(db)
}
