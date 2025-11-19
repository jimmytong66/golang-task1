package assignment4

import (
	"gorm.io/gorm"
)

// 题目 3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量统计字段
	return tx.Debug().Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 检查文章的评论数量
	var commentCount int64
	if err := tx.Debug().Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}
	// 如果评论数量为 0，则更新文章的评论状态为 "无评论"
	if commentCount == 0 {
		if err := tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error; err != nil {
			return err
		}
	}
	return nil
}
