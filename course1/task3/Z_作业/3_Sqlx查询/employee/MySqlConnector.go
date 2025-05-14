package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var db *sqlx.DB

func initDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Hour)

	log.Println("初始化数据库连接成功")
}

func closeDB() {
	err := db.Close()
	if err != nil {
		log.Println("Close db error:", err)
	}
	log.Println("关闭数据库连接成功")
}
