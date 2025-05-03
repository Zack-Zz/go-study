package main

import "fmt"

func main() {
	var nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob_0(nums))

	fmt.Println(rob_1(nums))
}

// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/
// 给定一个整数数组 nums ，代表每个房屋的金额
// 返回在不触动相邻房屋的情况下，能够偷窃到的最高金额。
// 1. dp[i] = max(dp[i-1], dp[i-2] + nums[i])
// 2. dp[0] = nums[0]
// 3. dp[1] = max(nums[0], nums[1])
// 4. dp[i] = max(dp[i-1], dp[i-2] + nums[i])

func rob_0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// 使用动态的数组，维护first和second两个变量
func rob_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		temp := second
		second = max(second, first+nums[i])
		first = temp
	}

	return second
}
