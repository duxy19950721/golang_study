package main

import (
	"golang_study/gorm/homework/sql_test/config"

	"gorm.io/gorm"
)

func main() {
	db := config.InitDB()

	db.AutoMigrate(&User{}, &Company{})
}

// 原始结构
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company

	// CompanyRefer string  `gorm:"size:100"` // 外键字段名和类型都改了
	// Company      Company `gorm:"foreignKey:CompanyRefer;references:Code"`
}

type Company struct {
	gorm.Model
	Name string
	// Code string `gorm:"unique"` // 新增唯一字段
}

// 修改后的结构
// type UserV2 struct {
// 	gorm.Model
// 	Name         string
// 	CompanyRefer string  // 外键字段名和类型都改了
// 	Company      Company `gorm:"foreignKey:CompanyRefer;references:Code"`
// }

// type CompanyV2 struct {
// 	gorm.Model
// 	Code string `gorm:"unique"` // 新增唯一字段
// 	Name string
// }
