package main

import (
	"fmt"
	datastruct "test/test/data_struct/normal"
)

// 匿名体结构
var b = struct {
	Age  int
	Name string
}{
	Age:  18,
	Name: "张三",
}

func main() {
	// 普通类型

	// 带匿名字段的类型
	var a datastruct.AnonymousStruct

	a.Age = 18
	a.Name = "张三"
	a.SetAnonymous("男")
	fmt.Println(a)

	// 匿名体结构
	c := struct {
		Age  int
		Name string
	}{
		Age:  18,
		Name: "张三",
	}
	fmt.Println(c)

}
