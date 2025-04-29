package main

import "fmt"

// 方式1
const a int = 1

// 方式2
const b = "test"

// 方式3
const c, d = 2, "hello"

// 方式4
const e, f bool = true, false

// 方式5
const (
	h    byte = 3
	i         = "value"
	j, k      = "v", 4
	l, m      = 5, false
)

const (
	n = 6
)

// 枚举方式1
const (
	Male   = "Male"
	Female = "Female"
)

// Gender 枚举方式2
type Gender string

const (
	Male1   Gender = "Male1"
	Female1 Gender = "Female1"
)

func printGender(gender Gender) {
	fmt.Println(gender)
}

func (g *Gender) getGender() string {
	switch *g {
	case Male1:
		return "Male2"
	case Female1:
		return "Female2"
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male
}

func main() {
	fmt.Println(Male)
	fmt.Println(Female)

	printGender(Male1)

	male1 := Male1
	fmt.Println(male1.getGender())

}
