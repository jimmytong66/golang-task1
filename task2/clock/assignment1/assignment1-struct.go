// package main

// import (
// 	"fmt"
// 	"sync"
// )

// /*
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。 */

// type Counter struct {
// 	value int        //计数值
// 	mu    sync.Mutex //同步锁
// }

// // 自增方法
// func (c *Counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// func (c *Counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value

// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := Counter{}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			for j := 0; j < 1000; j++ {
// 				counter.increment()
// 			}
// 			fmt.Printf("Goroutine %d completed\n", id)
// 		}(i)
// 	}
// 	wg.Wait()
// 	// 输出最终结果
// 	fmt.Printf("Final counter value: %d\n", counter.getValue())
// 	fmt.Printf("Expected value: %d\n", 10*1000)
// }
