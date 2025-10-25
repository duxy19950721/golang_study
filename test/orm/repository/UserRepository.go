package repository

import (
	"fmt"
	"golang_study/test/orm/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (res *UserRepository) GetUser() {
	var user model.User
	res.Db.Last(&user)
	fmt.Println("查询成功", user)
}

func (res *UserRepository) Insert(user *model.User) {
	result := res.Db.Create(user)
	fmt.Println("插入成功", user.Model.ID, result)
}

func (res *UserRepository) InsertTargetField(user *model.User) {
	res.Db.Select("Birthday").Create(user)
}
