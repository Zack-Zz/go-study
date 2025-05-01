package main

import (
	"fmt"
	"time"
)

func main() {
	defer traceCost(time.Now(), "main")

	a := A{1}
	// 把方法赋值给函数变量
	function1 = a.add

	// 声明一个闭包并直接执行
	// 此闭包返回值是另外一个闭包（带参闭包）
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is a anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "test")
	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

	fmt.Println("a.i = ", a.i)
	fmt.Println("after call function1, a.i = ", function1(1))
	fmt.Println("a.i = ", a.i)

	var pf = func(i int) int {
		i++
		// 设置耗时
		time.Sleep(3 * time.Second)
		return i
	}
	incr_s := memoize(pf)

	start := time.Now()
	fmt.Println("incr_s fir:", incr_s(1))
	fmt.Printf("耗时: %.2f s\n", time.Since(start).Seconds())

	start1 := time.Now()
	fmt.Println("incr_s sec:", incr_s(1))
	fmt.Printf("耗时: %.2f s\n", time.Since(start1).Seconds())
}

// 闭包实现的简易缓存器
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		res := f(n)
		cache[n] = res
		return res
	}
}

// 斐波那契数列
func slowFib(n int) int {
	if n <= 1 {
		return n
	}
	time.Sleep(3 * time.Second) // 模拟耗时
	return slowFib(n-1) + slowFib(n-2)
}

// 跟踪耗时
func traceCost(start time.Time, name string) {
	cost := time.Since(start)
	fmt.Printf("[%s] 耗时: %v\n", name, cost)
}

/** 函数
func <function_name>(<parameter list>) (<return types>) {
    <expressions>
    ...
}
*/

/** 闭包
// 声明函数变量
var <closure name> func(<parameter list>) (<return types>)

// 声明闭包
var  <closure name> func(<parameter list>) (<return types>) = func(<parameter list>) (<return types>) {
    <expressions>
    ...
}

// 声明并立刻执行
func(<parameter list>) (<return types>) {
    <expressions>
    ...
}(<value list>)

// 作为参数，并调用
func <function name>(...,<name> func(<parameter list>) (<return types>), ...) {
    ...
    <var1>,... := <name>(<value list>)
    ...
}

// 作为返回值
func <function name>(...) (func(<parameter list>) (<return types>)) {
    ...
    <var1>,... := <name>(<value list>)
    ...
}
*/

type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var function1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}
