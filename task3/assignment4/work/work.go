package work

/*
## 进阶 gorm

进阶 gorm
题目 1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用 Gorm 定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写 Go 代码，使用 Gorm 创建这些模型对应的数据库表。
题目 2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写 Go 代码，使用 Gorm 查询某个用户发布的所有文章及其对应的评论信息。
编写 Go 代码，使用 Gorm 查询评论数量最多的文章信息。
题目 3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。*/

import (
	"fmt"

	"github.com/jimmytong/golang-task/task3/assignment4"
	"gorm.io/gorm"
)

// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func GetUserPostsAndComments(db *gorm.DB, userID uint) (assignment4.User, error) {
	var user assignment4.User
	result := db.Debug().Preload("Posts.Comments").First(&user, userID)
	if result.Error != nil {
		return user, result.Error
	}
	fmt.Printf("用户: %s\n", user.Name)
	return user, result.Error
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func GetMostCommentedPost(db *gorm.DB) (assignment4.Post, error) {
	var post assignment4.Post
	// 使用 Joins 关联查询，并按评论数降序排序，取第一条
	if err := db.Debug().Preload("User").Preload("Comments").Joins("JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("COUNT(comments.id) DESC").
		First(&post).Error; err != nil {
		return assignment4.Post{}, err
	}
	return post, nil
}
