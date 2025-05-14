package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func initDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}

	log.Println("初始化数据库连接成功")
}

func closeDB() {
	hsqldb, err := DB.DB()
	if err != nil {
		log.Println("Close db error:", err)
	}
	err1 := hsqldb.Close()
	if err1 != nil {
		log.Println("Close db error:", err1)
	}
	log.Println("关闭数据库连接成功")
}
