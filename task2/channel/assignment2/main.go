package main

import (
	"fmt"
	"sync"
)

/* 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。 */

func produce(ch chan int, wg *sync.WaitGroup) {
	// wg.Add(1)
	defer close(ch)
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	fmt.Println("所有数据已生产") // 可选：增加执行完成的提示
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		fmt.Printf("%d\n", c)
	}
	fmt.Println("所有数据已消费") // 可选：增加执行完成的提示
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	wg.Add(2)
	go produce(ch, &wg)
	go consume(ch, &wg)
	wg.Wait()
}
