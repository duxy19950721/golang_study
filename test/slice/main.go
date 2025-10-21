package main

import "fmt"

func main() {
	//test1()
	//test2()
	//test3()
	// add()
	// remove()
	testCopy()
}

func test1() {
	a := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	a[0] = 9
	a[1] = 8
	a[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
}

func test2() {
	s1 := []int{5, 4, 3, 2, 1}
	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(s1)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

	// 向指定位置赋值
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Println(s1)

	// range迭代访问切片
	for i, v := range s1 {
		fmt.Println("before modify, s1[%d] = %d", i, v)
	}
}

func test3() {
	s3 := []int{}
	fmt.Println("s3 = ", s3)

	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println("s3 = ", s3)
}

func add() {
	s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Println("s4 = ", s4)
}

func remove() {
	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)
}

func testCopy() {
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}
