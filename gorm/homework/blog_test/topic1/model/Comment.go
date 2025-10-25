package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("Comment AfterDelete comments:", comment)
	result := tx.Model(&Comment{}).Where("post_id = ?", comment.PostID).Find(&Comment{})
	fmt.Println("Comment AfterDelete result:", result)

	if result.RowsAffected == 0 {
		result := tx.Model(&Post{}).Where("id = ?", comment.PostID).Update("comments_status", "无评论")
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := tx.Model(&Post{}).Where("id = ?", comment.PostID).Update("comments_status", "有评论")
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
