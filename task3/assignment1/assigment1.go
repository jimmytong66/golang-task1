package assignment1

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/* 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。 */

type Student struct {
	ID   uint
	Name string
	// Info Info `gorm:"embedded"`
	Age   uint8
	Grade string
}

// type Info struct {
// 	Gender string
// 	Email string
// }

func GenericsAPIRun(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	ctx := context.Background()
	// studentPtr := &Student{
	// 	Name: "张三",
	// 	Age: 22,
	// 	Grade: "三年级",
	// }
	// result := gorm.WithResult()
	// err := gorm.G[Student](db, result).Create(ctx, studentPtr)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(studentPtr.ID)
	// fmt.Println(result.RowsAffected)

	studentsPtr := []*Student{
		{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"},
	}
	// studentsPtr := []*Student{
	// 	&Student{Name: "jinzhu1"}, &Student{Name: "jinzhu2"}, &Student{Name: "jinzhu3"},
	// }

	// students := []Student{
	// 	{Name: "jinzhu111"}, {Name: "jinzhu222"}, {Name: "jinzhu333"},
	// }

	results := gorm.WithResult()
	// gorm.G[[]Student](db, results).Create(ctx, &students)
	gorm.G[[]*Student](db, results).Create(ctx, &studentsPtr)
	fmt.Println(studentsPtr)
}

func Field(db *gorm.DB) {
	// 声明切片变量用于接收查询结果
	students := []Student{}
	db.Debug().Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&students, []int{3, 2, 1})
	fmt.Println(students)
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Student{})

}

func CreateStudents(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	result := db.Debug().Create(&student)
	fmt.Println(student.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

func QueryStudents(db *gorm.DB) {
	students := []Student{}
	db.Debug().Where("age > ?", "18").Find(&students)
	fmt.Println(students)
}

func UpdateGrade(db *gorm.DB) {

	// db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	db.Debug().Model(&Student{}).Where("name = ?", "张三").Updates(&Student{Grade: "四年级"})
}

func DeleteStudents(db *gorm.DB) {
	db.Debug().Where("age < ?", 15).Delete(&Student{})
}
