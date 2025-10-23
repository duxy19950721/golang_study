package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	for index := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || strs[0][:index] != strs[i][:index] {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}
