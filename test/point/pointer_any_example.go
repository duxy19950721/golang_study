package main

import "fmt"

// 初始化指针对象
var a *int

// 初始化一个变量对象
var b int = 10

func main() {
	// 给指针对象进行赋值
	a = &b
	// 打印指针
	fmt.Printf("a = %p, *a = %d\n", a, *a)

	// 再次改变指针的引用对象
	c := 20
	a = &c
	fmt.Printf("a = %p, *a = %d\n", a, *a)

	// 进阶 any
	var d any
	e := "字符串"
	d = &e
	if p, ok := d.(*string); ok {
		fmt.Printf("d是字符串类型: %s\n", p)
	}

	*a = 100
	fmt.Printf("a = %p, *a = %d\n", a, *a)
}

// func main() {
// 	// 1. 指针指向具体类型（如int）- 这是合法的
// 	var num int = 42
// 	var ptr *int = &num
// 	fmt.Println("指针指向int类型:")
// 	fmt.Printf("num = %d\n", num)
// 	fmt.Printf("ptr = %p, *ptr = %d\n\n", ptr, *ptr)

// 	// 2. 直接使用any类型 - 这是合法的
// 	// any是interface{}的别名，本身就是引用类型
// 	var value any = "hello"
// 	fmt.Println("使用any类型:")
// 	fmt.Printf("value = %v, type = %T\n\n", value, value)

// 	// 3. 使用指针指向具体类型，然后赋值给any - 这也是合法的
// 	var ptrValue any = ptr
// 	fmt.Println("any类型存储指针:")
// 	fmt.Printf("ptrValue = %v\n", ptrValue)
// 	if p, ok := ptrValue.(*int); ok {
// 		fmt.Printf("从any中获取指针并解引用: %d\n\n", *p)
// 	}

// 	// 4. 以下注释的代码是非法的，取消注释会导致编译错误
// 	// var invalid *any  // 错误: cannot use type *any as type *interface{}
// 	// var invalid2 *interface{}  // 错误: cannot use type *interface{} as type *interface{}

// 	/*
// 		解释为什么 *any 或 *interface{} 是不允许的：
// 		1. 在Go中，接口(interface)本身已经是引用类型，通常不需要再使用指针指向接口
// 		2. 接口变量存储的是值的类型信息和数据指针，已经包含了引用语义
// 		3. Go语言规范不允许直接声明指向接口类型的指针
// 		4. 如果需要存储接口类型的地址，应该使用reflect包中的ValueOf函数
// 	*/

// 	// 5. 正确的做法：如果需要存储接口变量的地址，使用reflect包
// 	fmt.Println("使用reflect包处理接口地址:")
// 	fmt.Println("(请参考reflect包文档了解更多细节)")
// }
