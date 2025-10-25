package main

import (
	"fmt"
	"golang_study/test/orm/config"
	"golang_study/test/orm/model"
	"golang_study/test/orm/repository"
	"time"
)

func main() {
	db := config.InitDB()
	fmt.Println("数据库连接成功:", db)

	db.AutoMigrate(&model.User{})
	fmt.Println("数据库表迁移成功")

	userRepository := repository.UserRepository{Db: db}

	// 方法1: 先创建时间变量，再取地址
	now := time.Now()
	user := model.User{Age: 18, Birthday: &now}
	userRepository.InsertTargetField(&user)
	userRepository.GetUser()
}
