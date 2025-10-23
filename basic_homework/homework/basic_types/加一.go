package main

import "fmt"

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i]%10 != 0 {
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
