package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

var taskTimeMap map[int]int = make(map[int]int)

func main() {
	tasks := make([]Task, 0, 10)
	for i := 1; i <= 10; i++ {
		tasks = append(tasks, func() {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		})
	}
	scheduler(tasks)
}

type Task func()

// 任务调度器
func scheduler(tasks []Task) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for i := range tasks {
		// 这里要注意，闭包里面的i必须要传进来，或者从for循环内部copy一个，不然有问题
		go func(i int) {
			defer wg.Done()
			start := time.Now()
			tasks[i]()
			execution := time.Since(start)
			taskTimeMap[i] = int(execution.Seconds())
		}(i)
	}
	wg.Wait()
	fmt.Println(taskTimeMap)
}
