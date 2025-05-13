package main

import (
	"fmt"
	"sync"
)

/*
* 使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
 */
func main() {
	var wg = sync.WaitGroup{}
	wg.Add(10)

	var num int = 0

	lock := sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			lock.Lock()
			for j := 0; j < 1000; j++ {
				num++
			}
			lock.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}
