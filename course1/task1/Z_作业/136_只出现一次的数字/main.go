package main

import "fmt"

func main() {
	var nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber_1(nums))
	fmt.Println(singleNumber_2(nums))
}

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/

// 给定一个整数数组 nums ，除某个元素只出现一次以外，其余每个元素均出现两次。
// 找出并返回那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法，且不使用额外空间来解决此问题。

// 异或运算
// 1. 任何数和 0 做异或运算，结果仍然是原来的数。
// 2. 任何数和其自身做异或运算，结果是 0。
// 3. 异或运算满足交换律和结合律。
// 4. 如果 a^b^c = 0，那么 a^b = c，a^c = b，b^c = a。
func singleNumber_2(nums []int) int {
	var single int
	for _, value := range nums {
		// 5. 将所有数字进行异或运算，最后得到的结果就是只出现一次的数字
		single ^= value
	}
	return single
}

// 暴力删除
func singleNumber_1(nums []int) int {
	var nums2 []int
	// 1. 遍历数组，统计每个数字出现的次数
	for _, value := range nums {
		// 2. 如果数字已经存在于 nums2 中，则删除它
		exist := false
		for i, v := range nums2 {
			if v == value {
				nums2 = append(nums2[:i], nums2[i+1:]...)
				exist = true
				break
			}
		}
		// 3. 如果数字不存在于 nums2 中，则添加它
		if !exist {
			nums2 = append(nums2, value)
		}
	}

	// 4. 如果 nums2 中只有一个数字，则返回它
	return nums2[0]
}
