package main

import "log"

func main() {

	initDb()

	// 创建数据表
	err := DB.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Println("AutoMigrate error. ", err)
	}

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息

	//使用Gorm查询评论数量最多的文章信息

	closeDB()
}
