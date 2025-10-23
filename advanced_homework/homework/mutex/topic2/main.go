package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全

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
	count atomic.Int64
}

func (c *Counter) Increment() {
	c.count.Add(1)
}
