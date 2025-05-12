package main

import "fmt"

func main() {

	var nums []int = []int{9}

	fmt.Println("plus 1. nums = ", plusOne(nums))

}

func plusOne(digits []int) []int {
	n := len(digits)

	flag := false
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || flag {
			digits[i] = digits[i] + 1
		}

		if digits[i] == 10 {
			digits[i] = 0
			flag = true
		} else {
			flag = false
		}
	}

	if !flag {
		return digits
	}

	x := make([]int, 1)
	x[0] = 1
	x = append(x, digits...)
	return x
}
