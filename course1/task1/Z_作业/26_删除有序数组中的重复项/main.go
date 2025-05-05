package main

import "fmt"

func main() {
	var nums []int = make([]int, 0)
	p := []int{1, 1, 2, 3, 3, 4, 5}
	nums = append(nums, p...)

	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

// 26. 删除有序数组中的重复项
// 题目：给你一个升序排列的数组 nums ，请你原地 删除重复出现的元素，使每个元素 只出现一次 ，
// 返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	result := 1
	for idx := 1; idx < n; idx++ {
		if nums[idx] != nums[idx-1] {
			nums[result] = nums[idx]
			result++
		}
	}

	return result

}

func removeDuplicates_0(nums []int) int {
	var dup []int

	for _, value := range nums {
		if contains(dup, value) {
			continue
		}
		dup = append(dup, value)
	}

	nums = dup
	return len(dup)
}

func contains(nums []int, num int) bool {
	for _, value := range nums {
		if value == num {
			return true
		}
	}
	return false
}
