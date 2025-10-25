package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID        uint
	UserName      string
	Posts         []Post
	TotalPostsCnt int
}
