package main

import "fmt"

// func main() {
// 	shengming1()
// 	shengming2()
// 	shengming3()
// 	shengming4()
// }

func shengming1() {
	// 仅声明
	var arr1 [3]int
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	fmt.Println(arr1)
}

func shengming2() {
	// 声明类型+初始化
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)
}

func shengming3() {
	// 声明无类型+初始化
	var arr1 = []int{10, 20, 30}
	fmt.Println(arr1)
}

func shengming4() {
	// 声明无类型+初始化
	var arr1 []string = []string{0: "1111", 3: "2222"}
	fmt.Println(arr1)
}
