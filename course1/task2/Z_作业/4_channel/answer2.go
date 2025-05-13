package main

import (
	"fmt"
	"sync"
)

/*
 * 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
 */
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 带缓冲的chan， 50的缓冲长度，意味着写入50个数据都不会阻塞，超过50则会阻塞直到读取完成才会继续写入
	channel := make(chan int, 50)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	go func() {
		defer wg.Done()

		for v := range channel {
			fmt.Println("received data: ", v)
		}
	}()

	wg.Wait()
}
