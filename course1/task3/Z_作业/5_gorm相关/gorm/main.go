package main

import (
	"log"
	"time"
)

func main() {

	initDb()

	// 创建数据表
	err := DB.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Println("AutoMigrate error. ", err)
	}

	// 创建数据
	//createData()

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息

	// select p.* from users u left join posts p on u.id = p.user_id where u.Name = ?;
	var users []User
	DB.Preload("Posts").Preload("Comments").Where("id=?", "2").Find(&users)
	for _, u := range users {
		log.Println("Comments", u.Comments)
		log.Println("Posts", u.Posts)
	}

	log.Println("使用Gorm查询评论数量最多的文章信息")

	//使用Gorm查询评论数量最多的文章信息
	log.Println("使用Gorm查询评论数量最多的文章信息")
	var posts []Result
	DB.Model(&Post{}).Select("posts.*, count(comments.id) as comment_count").
		Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Scan(&posts)
	for _, post := range posts {
		log.Println("post comment_count: ", post.CommentCount)
	}

	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
	//createPost()

	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	deleteComment()

	//closeDB()

	time.Sleep(10 * time.Second)
}

type Result struct {
	Post
	CommentCount int64
}

func createData() {
	user := User{
		Name:       "小明",
		Age:        12,
		LastUpdate: time.Now(),
		Posts: []Post{
			{
				Title:      "小明的第一篇文章",
				Content:    "小明的第一篇文章内容",
				LastUpdate: time.Now(),
			},
			{
				Title:      "小明的第二篇文章",
				Content:    "小明的第二篇文章内容",
				LastUpdate: time.Now(),
			},
		},
		Comments: []Comment{},
	}

	DB.Create(&user)

	user2 := User{
		Name:       "小王",
		Age:        15,
		LastUpdate: time.Now(),
		Posts: []Post{
			{
				Title:      "小王的第一篇文章",
				Content:    "小王的第一篇文章内容",
				LastUpdate: time.Now(),
			},
			{
				Title:      "小王的第二篇文章",
				Content:    "小王的第二篇文章内容",
				LastUpdate: time.Now(),
			},
		},
		Comments: []Comment{
			{
				PostId:     user.Posts[0].Id,
				Content:    "小明的第一篇文章的评论",
				LastUpdate: time.Now(),
			},
		},
	}

	DB.Create(&user2)
}

func createPost() {
	DB.Create(&Post{
		UserId:     3,
		Title:      "小明的第三篇文章",
		Content:    "小明的第三篇文章内容",
		LastUpdate: time.Now(),
	})
}

func deleteComment() {
	var comment Comment
	var commentId int64 = 3
	if err := DB.First(&comment, commentId).Error; err != nil {
		log.Println("Get comment error. ", err)
		return
	}

	result := DB.Delete(&comment, commentId)
	if result.Error != nil {
		log.Println("Delete comment error. ", result.Error)
	} else {
		log.Println("Delete comment success.")
	}
	log.Println("rows affected:", result.RowsAffected)
}
