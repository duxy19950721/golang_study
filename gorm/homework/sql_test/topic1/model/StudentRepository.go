package model

import (
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

// 插入
func (sr *StudentRepository) Insert(student *Student) {
	result := sr.DB.Create(student)
	fmt.Println("插入成功", student.Model.ID, result)
}

// 查询大于18岁的学生信息
func (sr *StudentRepository) GetStudentThan18() []Student {
	students := make([]Student, 10)
	sr.DB.Where("age > ?", 18).Find(&students)
	fmt.Println("查询结果：", students)
	return students
}

// 更新name的grade=四年级
func (sr *StudentRepository) UpdateByName(name string) {
	result := sr.DB.Model(&Student{}).Where("name = ?", name).Update("grade", "四年级")
	fmt.Println("更新成功,条数：", result.RowsAffected, result.Error)
}

// 删除年龄<15的记录
func (sr *StudentRepository) DeleteByAge(age int) {
	result := sr.DB.Where("age < ?", age).Delete(&Student{})
	fmt.Println("删除成功,条数：", result.RowsAffected)
}
