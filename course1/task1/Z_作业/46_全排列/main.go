package main

import "fmt"

func main() {
	// leetcode 46. 全排列
	// test case
	nums := []int{1, 2, 3}

	result := permute(nums)

	// print result
	for _, v := range result {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Println()
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))

	var backTrack func()
	backTrack = func() {
		// define recursion terminated
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 遍历所有选择
		for index, value := range nums {
			if used[index] {
				// 跳出
				continue
			}

			// 做出选择
			path = append(path, value)
			used[index] = true

			// 进入下一层
			backTrack()

			// 回溯
			path = path[:len(path)-1]
			used[index] = false
		}

	}

	backTrack()

	return result

	//var result [][]int
	//var path []int
	//used := make([]bool, len(nums))
	//backtrack(nums, &result, path, used)
	//return result
}

func backtrack(nums []int, i *[][]int, path []int, used []bool) {
	if len(path) == len(nums) {
		// 复制path的值到result中
		var temp []int
		for _, v := range path {
			temp = append(temp, v)
		}
		*i = append(*i, temp)
		return
	}

	for j := 0; j < len(nums); j++ {
		if used[j] {
			continue
		}
		path = append(path, nums[j])
		used[j] = true
		backtrack(nums, i, path, used)
		path = path[:len(path)-1]
		used[j] = false
	}
}
