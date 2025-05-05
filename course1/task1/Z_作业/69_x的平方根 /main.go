package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(mySqrt(4))

}

// 69. x 的平方根
// 题目：给你一个非负整数 x ，计算并返回 x 的 平方根 。
// 由于返回类型是整数，结果只保留 整数 部分，
// 也就是说，我们只保留 小数点前的部分。
// 不能使用任何内置指数函数和算符，比如 pow(x, 0.5) 或者 x ** 0.5 。

// 使用自然数计算幂次方
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 任何数目的幂次方（n表示次方）都等于自然数e的（n*Inx）次方
	ans := int(math.Exp(0.5 * math.Log(float64(x))))

	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}

	return ans
}
