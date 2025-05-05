package main

import (
	"fmt"
	"sort"
)

func main() {

	var nums = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println(merge(nums))
}

func merge(intervals [][]int) [][]int {
	// sort the intervals by the first element
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 排序之后，无论如何，都会存在区间
	results := [][]int{intervals[0]}
	for _, nums := range intervals {
		// 取出最后一个区间,因为已经排序了，所以最后一个区间就是最大的
		last := results[len(results)-1]
		// 取出当前区间的结束值
		end := nums[1]
		// 取出结果集最后一个区间的结束值
		var r1 = last[1]
		if r1 >= nums[0] {
			if r1 < end {
				// 只需要修改结果集最后区间结束值
				last[1] = end
			}
		} else {
			// last[1] < nums[0]，当前区间的开始值大于结果集最后一个区间的结束值，则认为结果集需要添加一个新的区间
			results = append(results, nums)
		}
	}

	return results
}

func merge0(intervals [][]int) [][]int {
	var results [][]int

	// sort the intervals by the first element

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, nums := range intervals {
		pickArray(&results, nums)
	}

	return results
}

func pickArray(results *[][]int, nums []int) {
	found := false
	start, end := nums[0], nums[1]
	for _, res := range *results {
		if res[0] <= start && res[1] >= end {
			found = true
		}

		if res[0] >= start && res[0] <= end {
			res[0] = start
			found = true
		}
		if res[1] >= start && res[1] <= end {
			res[1] = end
			found = true
		}
	}
	if !found {
		*results = append(*results, nums)
	}
}
