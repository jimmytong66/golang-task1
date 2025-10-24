package main

import (
	"fmt"
	"sync"
)

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。 */
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&counter, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(counter)
}
// }

func increment(counterPtr *int, wgPtr *sync.WaitGroup, muPtr *sync.Mutex) {
	defer wgPtr.Done()
	for i := 0; i < 1000; i++ {
		muPtr.Lock()
		*counterPtr++
		muPtr.Unlock()
	}
}
