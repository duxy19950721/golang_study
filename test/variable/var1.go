package main

import "fmt"

var a string = "var1"

const b = "var2"

var c, d, e int = 1, 2, 3
var f bool

var (
	g int     = 4
	h bool    = true
	i float64 = 3.14
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}
