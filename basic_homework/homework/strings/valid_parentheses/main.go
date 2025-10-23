package main

import "fmt"

func main() {
	fmt.Println(isValid("()]"))
}

func isValid(s string) bool {
	var tmp = make([]rune, 1, 3)
	for _, s1 := range s {
		if s1 == '{' || s1 == '[' || s1 == '(' {
			tmp = append(tmp[:len(tmp)-1], s1)
		} else {
			if len(tmp) == 0 {
				return false
			}
			top := tmp[len(tmp)-1]
			if s1 == '}' && top == '{' || s1 == ']' && top == '[' || s1 == ')' && top == '(' {
				tmp = tmp[:len(tmp)-1]
			} else {
				return false
			}
			return len(tmp) == 0
		}
	}
	return false
}
