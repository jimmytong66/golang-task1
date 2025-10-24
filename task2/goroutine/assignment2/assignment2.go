package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

考察点 ：协程原理、并发任务调度。
*/
func taskManager(tasks []func(), wg *sync.WaitGroup) {
	for i, task := range tasks {
		wg.Add(1)
		go func(id int, task func()) {
			defer wg.Done()
			starttime := time.Now()
			task()
			costtime := time.Since(starttime)
			fmt.Println("执行第", i+1, "个任务的执行时间是: ", costtime)
		}(i, task)
	}

}

func main() {
	var wg sync.WaitGroup

	tasks := []func(){
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("任务3完成")
		},
	}

	taskManager(tasks, &wg)
	wg.Wait()

}
