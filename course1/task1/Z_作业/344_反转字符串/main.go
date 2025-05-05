package main

func main() {

}

// 循环反转
func reverseString(s []byte) {
	var b byte
	var length int = len(s)
	for idx, value := range s {
		if idx >= length/2 {
			break
		}
		b = s[length-1-idx]
		s[length-1-idx] = value
		s[idx] = b
	}
}
