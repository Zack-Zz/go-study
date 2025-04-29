package main

import "fmt"

func defineInt() {
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf

	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17

	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111

	// 十进制
	var h uint8 = 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}

func defineFloat() {
	var float1 float32 = 10
	float2 := 10.0

	// 会有编译错误：cannot use float2 (variable of type float64) as float32 value in assignment
	// float1 = float2

	float1 = float32(float2)

	fmt.Println(float1)
	fmt.Println(float2)
}

func defineComplex() {
	var z complex64 = 3 + 2i
	fmt.Println(z)

	// 复数类型和浮点型有类似的机制，默认的自动推导的复数类型是 complex128
	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1) // c2与c3是等价的

	x := real(c2)
	y := imag(c2)

	fmt.Println(c1)
	fmt.Println(c3)
	fmt.Println(x)
	fmt.Println(y)
}

func defineByte() {
	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)

	var s1 string = string(bytes)
	fmt.Println(s1)
}

func defineRune() {
	var r1 rune = 'a'
	var r2 rune = '世'

	// 字符串可以直接转换成 []rune（rune 切片）
	var s string = "abc，你好，世界！"
	var runes []rune = []rune(s)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(runes)
}
