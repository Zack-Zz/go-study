package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	doGoRoutine()

	counterSafeIncr()

	optChannel()
}

/**
goroutine 是轻量线程，创建一个 goroutine 所需的资源开销很小，所以可以创建非常多的 goroutine 来并发工作。

它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务分配给 CPU 执行的过程。

但是 goroutine 不是通常理解的线程，线程是操作系统调度的。

在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，交给 goroutine 执行即可。
*/

func doGoRoutine() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	go func(string) {
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/**
线程安全安全
在 Go 中，goroutine 是轻量级的线程，Go 运行时会自动调度 goroutine 的执行。
*/
// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

func counterSafeIncr() {
	counter := SafeCounter{}

	// 启动100个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second * 3)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())
}

/**
channel 是 Go 中定义的一种类型，专门用来在多个 goroutine 之间通信的线程安全的数据结构。

可以在一个 goroutine 中向一个 channel 中发送数据，从另外一个 goroutine 中接收数据。

channel 类似队列，满足先进先出原则。
*/
// 只接收channel的函数，参数为只读通道
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数，参数为只写通道
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func optChannel() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
