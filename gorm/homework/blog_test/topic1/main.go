package main

import (
	"fmt"
	"golang_study/gorm/homework/blog_test/topic1/model"
	"golang_study/gorm/homework/sql_test/config"

	"gorm.io/gorm"
)

func main() {

	db := config.InitDB()

	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	// 插入mock数据
	// createUser(db)
	createUserV2(db)

	// 查询评论数最多的文章，也就是comment最多的postId
	// getPostWithMostComments(db)

	// 删除评论，测试更新post的评论状态
	// deleteComment(db, 63)
}

func createUser(db *gorm.DB) {
	db.Create(&model.User{
		UserName: "张三",
		UserID:   1,
		Posts: []model.Post{
			{Title: "张三第1篇博客", Comments: []model.Comment{
				{Content: "张三第1篇博客的第1条评论"},
				{Content: "张三第1篇博客的第2条评论"},
				{Content: "张三第1篇博客的第3条评论"},
			}},
			{Title: "张三第2篇博客", Comments: []model.Comment{{Content: "张三第2篇博客的评论"}}},
			{Title: "张三第3篇博客", Comments: []model.Comment{{Content: "张三第3篇博客的评论"}}},
		},
	})

	db.Create(&model.User{
		UserName: "李四",
		UserID:   1,
		Posts: []model.Post{
			{Title: "李四第1篇博客", Comments: []model.Comment{
				{Content: "李四第1篇博客的第1条评论"},
			}},
			{Title: "李四第2篇博客", Comments: []model.Comment{
				{Content: "李四第2篇博客的评论"},
			}},
		},
	})
}

func createUserV2(db *gorm.DB) {
	user := model.User{
		UserName: "李四",
	}
	db.Create(&user)

	db.Create(&model.Post{
		Title:    "李四第1篇博客",
		Comments: []model.Comment{{Content: "李四第1篇博客的第1条评论"}},
		UserID:   user.ID,
	})

	db.Create(&model.Post{
		Title:    "李四第2篇博客",
		Comments: []model.Comment{{Content: "李四第2篇博客的第1条评论"}},
		UserID:   user.ID,
	})
}

func getPostWithMostComments(db *gorm.DB) {
	mostPost := model.Post{}

	// 第一种办法,子查询的方式
	subQuery := db.Model(&model.Comment{}).Group("post_id").Order("count(id) desc").Limit(1).Select("post_id")
	db.Preload("Comments").Where("id = (?)", subQuery).First(&mostPost)

	// 第二种办法，gorm的join
	// db.Preload("Comments").
	// 	Select("posts.id, count(comments.id)").
	// 	Joins("left join comments on posts.id = comments.post_id").
	// 	Group("posts.id").Order("count(posts.id) desc").
	// 	First(&mostPost)

	fmt.Println("评论数最多的文章信息是:", mostPost)
}

func deleteComment(db *gorm.DB, commentId uint) {
	comment := model.Comment{}
	db.Find(&comment, commentId)
	db.Delete(&comment)
}
