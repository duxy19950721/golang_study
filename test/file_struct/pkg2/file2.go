package pkg2

import (
	"fmt"
)

var a string = getName()

func getName() string {
	fmt.Println("pkg2 getName")
	return "pk2_name"
}

func init() {
	fmt.Println("file2 init")
}
