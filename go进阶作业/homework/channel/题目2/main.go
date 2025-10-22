package main

import (
	"fmt"
	"time"
)

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制

func main() {

	ch := make(chan int, 100)

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
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("send发送: %d\n", i)
	}
}

func receive(ch <-chan int) {
	for i := range ch {
		fmt.Printf("receive接收到: %d\n", i)
	}
}
