package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	employee := Employee{
		EmployeeId: 1,
		Person: Person{
			Name: "Wang",
			Age:  12,
		},
	}

	employee.PrintInfo()
}

/**
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者
*/

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Employee struct {
	Person
	EmployeeId int `json:"id"`
}

func (e Employee) PrintInfo() {
	// json.Marshal(e) 没有缩进
	jsonBytes, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Println("JSON 编码错误:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
