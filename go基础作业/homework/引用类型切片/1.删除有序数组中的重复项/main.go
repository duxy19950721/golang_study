package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	slow := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
