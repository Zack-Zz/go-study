package main

import "fmt"

func convertAny() {
	var s string = "Go语言"
	var bytes []byte = []byte(s)
	var runes []rune = []rune(s)

	fmt.Println("string length: ", len(s))
	fmt.Println("bytes length: ", len(bytes))
	fmt.Println("runes length: ", len(runes))
}

func convertAny2() {
	var s string = "Go语言"
	var bytes []byte = []byte(s)
	var runes []rune = []rune(s)

	fmt.Println("string sub: ", s[0:7])
	fmt.Println("bytes sub: ", string(bytes[0:7]))
	fmt.Println("runes sub: ", string(runes[0:3]))

	// 输出
	//string sub:  Go语�
	//bytes sub:  Go语�
	//runes sub:  Go语
}
