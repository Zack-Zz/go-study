package main

func main() {

}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x = x / 10
	}

	// 只有当x的位数为偶数时，x和revertNum相等。或者奇数时，x和revertNum/10相等
	return x == revertNum || x == revertNum/10
}
