// animal.go
package anamels

import "fmt"

type Animal struct{}

func (a Animal) Speak() { // 大写，可以被其他包调用
	fmt.Println("Animal speaks")
}

func (a Animal) whisper() { // 小写，只能animal包自己用
	fmt.Println("animal whispers")
}
