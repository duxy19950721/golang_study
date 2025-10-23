package main

import (
	"fmt"
	"orm/config"
)

func main() {
	db := config.InitDB()
	fmt.Println("数据库连接成功:", db)
	// userRepository := repository.UserRepository{Db: db}
	// userRepository.GetUser()
}
