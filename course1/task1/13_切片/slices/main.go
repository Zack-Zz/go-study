package main

import "fmt"

func main() {
	defineSlice2()

	useSlice()
}

func useSlice() {
	s1 := []int{5, 4, 3, 2, 1}
	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(s1, e1, e2, e3)

	// 向指定位置赋值
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Println(s1)

	// range迭代访问切片
	for i, v := range s1 {
		fmt.Println("before modify, s1[%d] = %d", i, v)
	}

	var nilSlice []int
	fmt.Println("nilSlice length:", len(nilSlice))
	fmt.Println("nilSlice capacity:", len(nilSlice))

	s2 := []int{9, 8, 7, 6, 5}
	fmt.Println("s2 length: ", len(s2))
	fmt.Println("s2 capacity: ", cap(s2))

	// 添加元素
	s3 := []int{}
	fmt.Println("s3 = ", s3)

	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println("s3 = ", s3)

	// 向指定位置添加元素
	s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Println("s4 = ", s4)

	// 指定位置移出
	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)

	// 复制切片
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("after copy, src1 = ", src1)
	fmt.Println("after copy, dst1 = ", dst1)

	fmt.Println("after copy, src2 = ", src2)
	fmt.Println("after copy, dst2 = ", dst2)
}

func defineSlice2() {
	a := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	a[0] = 9
	a[1] = 8
	a[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
}

func defineSlice() {

	// 方式1，声明并初始化一个空的切片
	var s1 []int = []int{}
	fmt.Println(s1)

	// 方式2，类型推导，并初始化一个空的切片
	var s2 = []int{}
	fmt.Println(s2)

	// 方式3，与方式2等价
	s3 := []int{}
	fmt.Println(s3)

	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	fmt.Println(s4)

	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	fmt.Println(s5)

	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)
	fmt.Println(s6)

	// 方式7，引用一个数组，初始化切片
	arr := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := arr[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := arr[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := arr[:2]

	fmt.Println(s7, s8, s9)
}
