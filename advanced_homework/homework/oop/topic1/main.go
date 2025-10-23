package main

import "fmt"

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

func main() {
	rec := Rectangle{A: 10}
	cir := Circle{A: 10}

	rec.Area()
	rec.Perimeter()
	cir.Area()
	cir.Perimeter()
}

type Shape interface {
	Area()
	Perimeter()
}

type Circle struct {
	A int
}

func (r *Circle) Area() {
	fmt.Println("yuan形的面积", r.A)
}
func (r *Circle) Perimeter() {
	fmt.Println("yuan形的周长", r.A)
}

type Rectangle struct {
	A int
}

func (r *Rectangle) Area() {
	fmt.Println("矩形的面积", r.A)
}
func (r *Rectangle) Perimeter() {
	fmt.Println("矩形的周长", r.A)
}
