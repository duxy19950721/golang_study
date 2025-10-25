package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title          string
	Comments       []Comment
	UserID         uint
	CommentsStatus string
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	result := tx.Model(&User{}).Where("id = ?", post.UserID).Update("total_posts_cnt", gorm.Expr("total_posts_cnt + 1"))
	fmt.Println("Post AfterCreate post, result:", post, result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
