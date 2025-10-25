package main

import (
	"fmt"
	"golang_study/gorm/homework/sql_test/config"
	"golang_study/gorm/homework/sqlx_test/topic2/model"
)

func main() {
	db := config.InitDB()

	// 查询大于50元的书籍
	books := make([]model.Book, 0)
	db.Table("books").Select("*").Where("price > ?", 50).Order("price desc").Limit(3).Scan(&books)
	fmt.Println("最畅销的书，价格前3名分别是:", books)
}
