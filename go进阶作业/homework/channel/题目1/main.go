package main

import (
	"fmt"
	"time"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

func main() {

	ch := make(chan int)

	go send(ch)
	go receive(ch)

	timeout := time.After(3 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Println("通信结束")
			return
		}
	}

}

func send(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("send发送: %d\n", i)
	}
}

func receive(ch <-chan int) {
	for i := range ch {
		fmt.Printf("receive接收到: %d\n", i)
	}
}
