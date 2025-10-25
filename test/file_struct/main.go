package main

import (
	"fmt"
)

var a string = getName()

func getName() string {
	return "main_name"
}

func main() {
	fmt.Println("Main program running")
	fmt.Println("main variable a =", a)
}
