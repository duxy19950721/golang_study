package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sleep等待执行，否则main线程关闭后其他线程还没完成，不打印日志
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(2 * time.Second)

	// 使用计数器

	var sync sync.WaitGroup
	for i := 0; i < 10; i++ {
		sync.Add(1)
		go func(i int) {
			defer sync.Done()
			fmt.Println(i)
			defer sync.Done()
		}(i)
	}

	// 等待所有goroutine完成
	sync.Wait()
}
