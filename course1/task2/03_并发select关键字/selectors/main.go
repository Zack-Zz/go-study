package main

import "fmt"

func main() {

	useSelector()

}

/**
select 语义是和 channel 绑定在一起使用的，select 可以实现从多个 channel 收发数据，可以使得一个 goroutine 就可以处理多个 channel 的通信。

语法上和 switch 类似，有 case 分支和 default 分支，只不过 select 的每个 case 后面跟的是 channel 的收发操作。
语法上和switch的一些区别：
1. select 关键字和后面的 { 之间，不能有表达式或者语句。
2. 每个 case 关键字后面跟的必须是 channel 的发送或者接收操作
3. 允许多个 case 分支使用相同的 channel，case 分支后的语句甚至可以重复
*/

func useSelector() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		case z := <-ch3:
			fmt.Printf("receive %d from channel 3\n", z)
		}
	}
}
