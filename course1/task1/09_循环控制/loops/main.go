package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	// 方式1
	for i := 0; i < 10; i++ {
		fmt.Println("方式1，第", i+1, "次循环")
	}

	// 方式2
	b := 1
	for b < 10 {
		fmt.Println("方式2，第", b, "次循环")
		b++
	}

	// 方式3，无限循环
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	var started bool
	var stopped atomic.Bool
	for {
		// 控制goruoutine只启动一次
		if !started {
			started = true
			// 子协程，等 context 取消后退出
			go func() {
				for {
					// select 语句不会“绑定”channel，它是一个语法结构，用来监听多个 channel 操作（发送/接收）是否准备好
					// select 把所有 case 中的通道操作注册成监听事件
					// runtime 会判断哪个 channel 当前是“就绪状态”
					//		如果有：立即执行那条 case 分支
					//		如果没有：阻塞等待
					// 如果多个都准备好了：随机选一个执行
					// 如果所有都没准备好，而且没有 default：整个 select 阻塞
					select {
					// 等待 context 被取消的通知
					case <-ctx.Done():
						fmt.Println("ctx done")

						// 通知主协程退出
						stopped.Store(true)
						return

					default:
						fmt.Println("working...")
						time.Sleep(500 * time.Millisecond)
					}
				}
			}()
		}
		// 当主协程为true，则退出循环
		if stopped.Load() {
			fmt.Println("main goroutine")
			break
		}
	}

	// 遍历数组
	var a [10]string
	a[0] = "Hello"
	for i := range a {
		fmt.Println("当前下标：", i)
	}
	for i, e := range a {
		fmt.Println("a[", i, "] = ", e)
	}

	// 遍历切片
	s := make([]string, 10)
	s[0] = "Hello"
	for i := range s {
		fmt.Println("当前下标：", i)
	}
	for i, e := range s {
		fmt.Println("s[", i, "] = ", e)
	}

	m := make(map[string]string)
	m["b"] = "Hello, b"
	m["a"] = "Hello, a"
	m["c"] = "Hello, c"
	for i := range m {
		fmt.Println("当前key：", i)
	}
	for k, v := range m {
		fmt.Println("m[", k, "] = ", v)
	}
}
