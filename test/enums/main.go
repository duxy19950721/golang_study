package main

import "fmt"

type Gender string

const (
	Male   Gender = "男"
	Female Gender = "女"
)

func main() {
	var gender Gender = Male
	fmt.Println(gender.getGender()) // 通过实例调用方法
	fmt.Println(gender.IsMale())
	fmt.Println(Male)
}

func (g *Gender) getGender() string {
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male
}
