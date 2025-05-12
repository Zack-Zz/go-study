package main

func main() {

}

// 20.有效的括号，括号配对，用栈处理(数组实现)
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	var pairs map[byte]byte = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte = []byte{}

	for i := 0; i < len(s); i++ {
		var x, ok = pairs[s[i]]
		if ok {
			if len(stack) == 0 || stack[len(stack)-1] != x {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}

	}

	return len(stack) == 0
}
