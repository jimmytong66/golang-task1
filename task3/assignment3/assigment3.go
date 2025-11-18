package assignment3

/*Sqlx入门
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
ß
*/

import (
	"gorm.io/gorm"
)

type Employee struct {
	ID         uint
	Name       string
	Department string
	Salary     float64
}

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

func getEmployeesByDepartment(db *gorm.DB, department string) ([]Employee, error) {
	var employees []Employee
	result := db.Where("department = ?", department).Find(&employees)
	return employees, result.Error
}

func getHighestSalaryEmployee(db *gorm.DB) (Employee, error) {
	var employee Employee
	result := db.Order("salary DESC").First(&employee)
	return employee, result.Error
}

func getBooksByPrice(db *gorm.DB, price float64) ([]Book, error) {
	var books []Book
	result := db.Where("price > ?", price).Find(&books)
	return books, result.Error
}
