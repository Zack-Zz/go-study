package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* 使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
 */
func main() {
	var channel chan int = make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			channel <- i
		}
		close(channel)
	}()

	go func() {
		defer wg.Done()

		/**
		方式1： 直接for range channel，会阻塞读取channel的数据，直到channel关闭
		*/
		//for v := range channel {
		//	fmt.Println("goroutine channel read done", v)
		//}

		/**
		方式2： for select case 阻塞读取
		*/
		timeout := time.After(5 * time.Second)

		for {
			select {
			case v, ok := <-channel:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				fmt.Println("received value", v)
			case <-timeout:
				fmt.Println("timeout")
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
