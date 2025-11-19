package assignment4

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 模型定义
type User struct {	
	gorm.Model
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount int    `gorm:"default:0"`
}

type Post struct {
	gorm.Model
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentStatus string    
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	Post    Post   `gorm:"foreignKey:PostID"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}

// MySQL 数据库连接配置
const (
	DB_USER     = "root"           // MySQL 用户名
	DB_PASSWORD = ""       // MySQL 密码
	DB_HOST     = "localhost"      // MySQL 主机
	DB_PORT     = "3306"           // MySQL 端口
	DB_NAME     = "blog_system"    // 数据库名
)

// 初始化 MySQL 数据库
func InitDB() *gorm.DB {
	// 构建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 减少日志输出
	})
	if err != nil {
		panic("failed to connect to MySQL database: " + err.Error())
	}

	// 自动迁移模式 - 创建表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	
	return db
}

// 生成用户数据
func createUser(db *gorm.DB, name, email string) User {
	user := User{
		Name:  name,
		Email: email,
	}
	db.Create(&user)
	return user
}

// 生成文章数据
func createPost(db *gorm.DB, title, content string, userID uint) Post {
	post := Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}
	db.Create(&post)
	return post
}

// 生成评论数据
func createComment(db *gorm.DB, content string, postID, userID uint) Comment {
	comment := Comment{
		Content: content,
		PostID:  postID,
		UserID:  userID,
	}
	db.Create(&comment)
	return comment
}

// 生成测试数据
func generateTestData(db *gorm.DB) {
	fmt.Println("开始生成 MySQL 测试数据...")
	
	// 创建用户
	users := []User{}
	userNames := []string{
		"张三", "李四", "王五", "赵六", "钱七", "孙八", "周九", "吴十",
		"郑一", "王二", "冯三", "陈四", "褚五", "卫六", "蒋七", "沈八",
	}
	
	for i, name := range userNames {
		email := fmt.Sprintf("user%d@example.com", i+1)
		user := createUser(db, name, email)
		users = append(users, user)
		fmt.Printf("创建用户: %s (%s)\n", user.Name, user.Email)
	}
	
	// 为每个用户创建文章
	postTitles := []string{
		"Go语言入门指南", "GORM使用技巧", "数据库设计最佳实践",
		"Web开发安全注意事项", "API设计原则", "微服务架构详解",
		"性能优化实战", "代码重构的艺术", "测试驱动开发",
		"前端框架选择", "DevOps实践", "容器化部署",
		"消息队列应用", "缓存策略", "分布式系统设计",
		"数据结构与算法", "设计模式应用", "版本控制最佳实践",
		"数据库索引优化", "API性能调优", "Go并发编程",
		"Redis应用实践", "Elasticsearch搜索优化", "Docker容器编排",
		"Kubernetes集群管理", "消息中间件选型", "API网关设计",
		"服务注册与发现", "链路追踪实践", "分布式锁实现",
	}
	
	// 为每个用户随机创建2-5篇文章
	for i, user := range users {
		articleCount := rand.Intn(4) + 2 // 每个用户2-5篇文章
		for j := 0; j < articleCount; j++ {
			titleIndex := (i*articleCount + j) % len(postTitles)
			title := postTitles[titleIndex]
			
			// 生成文章内容
			content := fmt.Sprintf(
				"这是关于%s的文章内容。Go语言是一门开源的编程语言，它能够让构造简单、可靠且高效的软件变得容易。\n\n"+
					"在本文中，我们将探讨%s的各种应用场景和技术细节。首先，我们需要了解其基本概念和原理。\n\n"+
					"通过实际代码示例，您将学会如何在项目中应用这一技术。最后，我们还会讨论一些高级特性和最佳实践。\n\n"+
					"希望这篇文章能对您有所帮助，如果您有任何问题或建议，请在评论区留言。",
				title, title)
			
			post := createPost(db, title, content, user.ID)
			fmt.Printf("  - 创建文章: %s\n", post.Title)
			
			// 为文章添加评论 (1-10条评论)
			commentCount := rand.Intn(10) + 1
			for k := 0; k < commentCount; k++ {
				// 随机选择评论者（可以是文章作者或其他用户）
				commenterIndex := rand.Intn(len(users))
				commenter := users[commenterIndex]
				
				commentContents := []string{
					"写得很好，学到了很多！",
					"感谢分享，很有帮助",
					"这个方法很实用",
					"期待更多类似的文章",
					"有更详细的代码示例吗？",
					"这个技术在实际项目中表现如何？",
					"能详细解释一下这部分吗？",
					"我有不同的看法...",
					"实际应用中需要注意什么？",
					"性能如何？",
					"和同类技术相比有什么优势？",
					"适合初学者吗？",
					"生产环境验证过吗？",
					"文档链接能提供一下吗？",
					"这个方案的缺点是什么？",
					"代码能贴出来看看吗？",
					"这个解决方案很巧妙",
					"在高并发场景下如何？",
					"有没有性能测试数据？",
					"这个库还在维护吗？",
				}
				
				commentContent := commentContents[rand.Intn(len(commentContents))]
				comment := createComment(db, commentContent, post.ID, commenter.ID)
				fmt.Printf("    - 添加评论: \"%s\" (评论者: %s)\n", 
					comment.Content, commenter.Name)
			}
		}
	}
	
	fmt.Println("MySQL 测试数据生成完成!")
}

// 统计信息
func printStats(db *gorm.DB) {
	fmt.Println("\n===== 数据统计 =====")
	
	var userCount, postCount, commentCount int64
	
	db.Model(&User{}).Count(&userCount)
	db.Model(&Post{}).Count(&postCount)
	db.Model(&Comment{}).Count(&commentCount)
	
	fmt.Printf("用户总数: %d\n", userCount)
	fmt.Printf("文章总数: %d\n", postCount)
	fmt.Printf("评论总数: %d\n", commentCount)
	fmt.Printf("平均每用户文章数: %.2f\n", float64(postCount)/float64(userCount))
	fmt.Printf("平均每文章评论数: %.2f\n", float64(commentCount)/float64(postCount))
	
	// 显示每个用户的文章数
	fmt.Println("\n每个用户的文章数:")
	var users []User
	db.Select("id, name, post_count").Find(&users)
	for _, user := range users {
		fmt.Printf("- %s: %d篇\n", user.Name, user.PostCount)
	}
}

func MockData() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	
	// 初始化数据库
	db := InitDB()
	
	// 生成测试数据
	generateTestData(db)
	
	// 打印统计信息
	printStats(db)
	
	fmt.Println("\n数据已保存到 MySQL 数据库")
	fmt.Println("数据库: blog_system")
	fmt.Println("表: users, posts, comments")
}