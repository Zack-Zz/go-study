package main

import "fmt"

func main() {

	num := 10
	addTen(&num)
	fmt.Println(num) // 20

	nums := []int{1, 2, 3, 4, 5}
	multiplyByTwo(nums)
	fmt.Println(nums)
}

// addTen 函数接收一个指向 int 的指针作为参数，并将其值加 10
func addTen(num *int) {
	*num = *num + 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func multiplyByTwo(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}
}
