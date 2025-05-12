package main

func main() {

}

/**
leetcode 1.两数之和
*/

// 暴力穷举,时间复杂度O(n^2)
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 哈希表,时间复杂度O(n)
func twoSum_0(nums []int, target int) []int {
	hashtable := map[int]int{}

	for i, x := range nums {
		// 本质上第二个循环用一个hash表代替，从而减少复杂度
		if p, ok := hashtable[target-x]; ok {
			return []int{i, p}
		}
		hashtable[x] = i
	}

	return []int{}
}
