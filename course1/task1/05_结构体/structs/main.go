package main

import (
	"fmt"
	"github.com/Zack-Zz/go-study/course1/structs/anamels"
)

func main() {
	a := anamels.Animal{}

	// 访问公开方法
	a.Speak()

	// 匿名结构
	PrintAnonymousStruct()

	p := Person{
		Name: "Tom",
		Address: Address{
			City:    "New York",
			Zipcode: "10001",
		},
	}

	fmt.Println(p.Address.City) // 直接访问

	fmt.Println(p.getPerson())
	fmt.Println(p.getOriginalPerson())
}
