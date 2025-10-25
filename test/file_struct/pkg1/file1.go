package pkg1

import (
	"fmt"
)

var a string = getName()

func getName() string {
	fmt.Println("pkg1 getName")
	return "pk1_name"
}

func init() {
	fmt.Println("file1 init")
}
