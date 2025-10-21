package main

import "fmt"

func main() {
	// 定义一个数组
	nums := []int{1, 2, 2}
	singleNumber(nums)
	// 输出结果
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var m = map[int]int{}
	for i := range nums {
		value := nums[i]
		if m[value] == 0 {
			m[value] = 1
		} else {
			m[value]++
		}
	}
	fmt.Println(m)
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
