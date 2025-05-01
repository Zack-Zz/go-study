package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	rangeStr()

	rangeArray()

	rangeMultiArray()

	rangeChan()

	rangeMappingCollection()
}

func rangeStr() {
	str1 := "abc123"
	for index := range str1 {
		fmt.Printf("str1 -- index:%d, value:%d, char:%c \n", index, str1[index], str1[index])
	}

	str2 := "测试中文"
	for index, ch := range str2 {
		fmt.Printf("str2 -- index:%d, value:%d, char:%c \n", index, str2[index], ch)
	}
	fmt.Printf("len(str2) = %d\n", len(str2))

	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))
}

func rangeArray() {
	array := [...]int{1, 2, 3}
	slice := []int{4, 5, 6}

	// 方法1：只拿到数组的下标索引
	for index := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	}
	for index := range slice {
		fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	}
	fmt.Println()

	// 方法2：同时拿到数组的下标索引和对应的值
	for index, value := range array {
		fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
		fmt.Printf("array -- index=%d range value=%d \n", index, value)
	}
	for index, value := range slice {
		fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
		fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	}
	fmt.Println()
}

func rangeMultiArray() {
	array := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slice := [][]int{{1, 2}, {3}}
	// 只拿到行的索引
	for index := range array {
		// array[index]类型是一维数组
		fmt.Println(reflect.TypeOf(array[index]))
		fmt.Printf("array -- index=%d, value=%v\n", index, array[index])
	}

	for index := range slice {
		// slice[index]类型是一维数组
		fmt.Println(reflect.TypeOf(slice[index]))
		fmt.Printf("slice -- index=%d, value=%v\n", index, slice[index])
	}

	// 拿到行索引和该行的数据
	fmt.Println("print array element")
	for row_index, row_value := range array {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	fmt.Println("print array slice")
	for row_index, row_value := range slice {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	// 双重遍历，拿到每个元素的值
	for row_index, row_value := range array {
		for col_index, col_value := range row_value {
			fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
	for row_index, row_value := range slice {
		for col_index, col_value := range row_value {
			fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
}

func rangeChan() {
	ch := make(chan int, 10)

	go addData(ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func rangeMappingCollection() {
	hash := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}

	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}
