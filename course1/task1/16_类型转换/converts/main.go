package main

import (
	"fmt"
	"strconv"
)

func main() {
	convertNumber()
	convertStr()

	convertInterface()

	convertConstruct()
}

func convertConstruct() {
	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	// 在字段类型、顺序、数量完全相同的两个 struct 之间做“显式类型转换”
	b := SameFieldB(a)
	fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", b.getValue())

	// 只能结构体类型实例之间相互转换，指针不可以相互转换
	var c interface{} = &a
	_, ok := c.(*SameFieldB)
	fmt.Printf("c is *SameFieldB: %v \n", ok)
}

type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (s *SameFieldB) getValue() int {
	return s.value
}

func convertInterface() {
	var i interface{} = 3
	a, ok := i.(int)
	if ok {
		fmt.Printf("'%d' is a int \n", a)
	} else {
		fmt.Println("conversion failed")
	}

	switch v := i.(type) {
	case int:
		fmt.Println("i is a int", v)
	case string:
		fmt.Println("i is a string", v)
	default:
		fmt.Println("i is unknown type", v)
	}
}

func convertNumber() {

	// 数字
	var i int32 = 17
	var b byte = 5
	var f float32

	// 数字类型可以直接强转
	f = float32(i) / float32(b)
	fmt.Printf("f 的值为: %f\n", f)

	// 当int32类型强转成byte时，高位被直接舍弃
	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("b2 的值为: %d\n", b2)
}

func convertStr() {
	/**
	字符串
	*/

	str := "hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes 的值为: %v \n", bytes)
	fmt.Printf("runes 的值为: %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 的值为: %v \n", str2)
	fmt.Printf("str3 的值为: %v \n", str3)

	// 字符串数字转换
	num, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	strO := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", strO)

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", num)

	strC := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", strC)
}
