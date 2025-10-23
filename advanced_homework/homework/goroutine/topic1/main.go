package main

import (
	"fmt"
	"time"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

func main() {
	ch := make(chan int, 3)

	go printOdd(ch)
	go printEven(ch)
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func printOdd(ch chan<- int) {
	for i := 1; i <= 10; i += 2 {
		ch <- i
		fmt.Println("已发送", i)
	}
}

func printEven(ch chan<- int) {
	for i := 2; i <= 10; i += 2 {
		ch <- i
		fmt.Println("已发送", i)
	}
}
