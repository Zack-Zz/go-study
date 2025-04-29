package main

import "fmt"

type Other struct{}

type Person struct {
	Name    string            `json:"name" gorm:"column:<name>"`
	Age     int               `json:"age" gorm:"column:<age>"`
	Call    func()            `json:"-" gorm:"column:<call>"`
	Map     map[string]string `json:"map" gorm:"column:<map>"`
	Ch      chan string       `json:"-" gorm:"-"`
	Arr     [32]uint8         `json:"arr" gorm:"column:<arr>"`
	Slice   []interface{}     `json:"slice" gorm:"column:<slice>"`
	Ptr     *int              `json:"-"`
	O       Other             `json:"-"`
	Address Address           `json:"address" gorm:"-"` // 嵌套了Address结构体
}

func PrintAnonymousStruct() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "Tom",
		Age:  20,
	}

	fmt.Println(person.Name, person.Age) // 输出: Tom 20
}

type Address struct {
	City    string
	Zipcode string
}

func (person Person) getPerson() Person {
	return Person{
		Name: "XiaoMing",
		Age:  22,
		Address: Address{
			City:    "SHai",
			Zipcode: "12312",
		},
	}
}

func (person Person) getOriginalPerson() Person {
	return person
}
