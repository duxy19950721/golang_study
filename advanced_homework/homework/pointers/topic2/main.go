package main

import "fmt"

//实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

func main() {
	var old_nums = []int{1, 2, 3}
	fmt.Println("old_nums:", old_nums)
	cheng2(&old_nums)
	fmt.Println("new_nums:", old_nums)
}

func cheng2(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}
