package main

import (
	"fmt"
	"golang_study/gorm/homework/sql_test/config"
	"golang_study/gorm/homework/sqlx_test/topic1/model"
)

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

func main() {
	db := config.InitDB()

	employees := make([]model.Employee, 0)
	// 查询部门为技术部的员工信息
	result := db.Table("employees").Select("*").Where("department = ?", "技术部")
	result.Scan(&employees)
	fmt.Println("技术部员工列表为:", employees)

	// 查询工资最高的员工信息
	high := model.Employee{}
	db.Table("employees").Select("*").Order("salary desc").First(&high)
	fmt.Println("工资最高的员工信息为:", high)
}
