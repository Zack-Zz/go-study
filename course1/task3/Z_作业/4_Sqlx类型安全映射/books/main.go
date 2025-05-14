package main

import (
	"log"
	"time"
)

func main() {

	initDb()

	log.Println(selectBook())

	closeDB()

}

type Book struct {
	ID         int64     `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Author     string    `json:"author" db:"author"`
	Price      float64   `json:"price" db:"price"`
	LastUpdate time.Time `json:"lastUpdate" db:"last_update"`
}

// 查询价格大于 30 元的书籍，并将结果映射到 Book 结构体切片中
func selectBook() []Book {
	var books []Book
	err := db.Select(&books, "select * from books where price > ?", 30)
	if err != nil {
		log.Fatalln("查询失败:", err)
	}
	return books
}
