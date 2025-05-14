package main

import (
	"log"
	"time"
)

func main() {
	initDb()

	selectEmployee()

	selectMaxSalaryEmployee()

	closeDB()
}

type Employee struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Department string    `json:"department" db:"department"`
	Salary     float64   `json:"salary" db:"salary"`
	LastUpdate time.Time `json:"lastUpdate" db:"last_update"`
}

// 使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
func selectEmployee() []Employee {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees where department = ?", "技术部")
	if err != nil {
		log.Fatalln("查询失败:", err)
	}
	for _, employee := range employees {
		log.Println("employee: ", employee)
	}
	return employees
}

// 使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
func selectMaxSalaryEmployee() Employee {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatalln("查询失败:", err)
	}
	log.Println("max salary employee: ", employee)
	return employee
}
