package main

import "fmt"

func defineStr() {

	// 一种是解释型字面表示（interpreted string literal，双引号风格）
	var s string = "Hello\nworld!\n"

	fmt.Println(s)

	// 另一种是直白字面量表示（raw string literal， 反引号风格）。
	var s1 string = `Hello
world!
`
	fmt.Println(s1 == s)
}
