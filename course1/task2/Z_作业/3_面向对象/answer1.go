package main

import (
	"fmt"
	"math"
)

/**
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。

然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。

在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
*/

func main() {
	var shape1 Shape = Rectangle{Width: 10, Height: 5}
	var shape2 Shape = Circle{Radius: 5}

	fmt.Println(shape1.Area())
	fmt.Println(shape1.Perimeter())
	fmt.Println(shape2.Area())
	fmt.Println(shape2.Perimeter())
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
