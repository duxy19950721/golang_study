package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

func main() {
	wg := sync.WaitGroup{}

	c := Counter{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {

			for i := 1; i <= 1000; i++ {
				c.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.count)
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
