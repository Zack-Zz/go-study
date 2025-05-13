package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//doGoroutinePrintNums()

	tasks := []func(){
		func() {
			time.Sleep(5 * time.Second)
			fmt.Println("hello world. 5s")
		}, func() {
			time.Sleep(3 * time.Second)
			fmt.Println("hello world. 3s")
		},
	}
	scheduleTask(tasks)
}

// 使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func doGoroutinePrintNums() {

	var wg sync.WaitGroup
	wg.Add(2)

	// 使用lock和条件队列实现两个goroutine交替打印数字
	lock := sync.Mutex{}
	cond := sync.NewCond(&lock)

	turn := 1

	go func() {
		defer wg.Done()
		for turn < 10 {
			lock.Lock()
			if turn%2 == 0 {
				cond.Wait()
				fmt.Println(turn)
			}
			turn++
			cond.Signal()
			lock.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for turn < 10 {
			lock.Lock()
			if turn%2 == 1 {
				cond.Wait()
				fmt.Println(turn)
			}
			turn++
			cond.Signal()
			lock.Unlock()
		}
	}()

	wg.Wait()
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
func scheduleTask(tasks []func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func() {
			defer wg.Done()
			defer timeConsume(time.Now())
			
			task()
		}()
	}

	wg.Wait()
}

func timeConsume(now time.Time) {
	fmt.Println("time consume:", time.Now().Sub(now))
}
