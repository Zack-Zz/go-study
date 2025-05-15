package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name       string    `json:"name" gorm:"not null;column:name"`
	Age        int       `json:"age" gorm:"not null;column:age"`
	PostCount  int64     `json:"post_count" gorm:"default:0;column:post_count"`
	LastUpdate time.Time `json:"last_update" gorm:"not null;column:last_update"`

	Posts    []Post    `json:"posts" gorm:"foreignKey:UserId;references:Id"`
	Comments []Comment `json:"comments" gorm:"foreignKey:UserId;references:Id"`
}

type Post struct {
	Id            int64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserId        int64     `json:"user_id" gorm:"not null;column:user_id;index"`
	Title         string    `json:"title" gorm:"not null;column:title"`
	Content       string    `json:"content" gorm:"not null;column:content"`
	LastUpdate    time.Time `json:"last_update" gorm:"not null;column:last_update"`
	CommentStatus string    `json:"comment_status" gorm:"default null;column:comment_status"`

	Comments []Comment `json:"comments" gorm:"foreignKey:PostId;references:Id"`
}

type Comment struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserId     int64     `json:"user_id" gorm:"not null;column:user_id;index"`
	PostId     int64     `json:"post_id" gorm:"not null;column:post_id;index"`
	Content    string    `json:"content" gorm:"not null;column:content"`
	LastUpdate time.Time `json:"last_update" gorm:"not null;column:last_update"`
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate hook triggered")
	return tx.Model(&User{}).
		Where("id = ?", p.UserId).
		UpdateColumn("post_count", gorm.Expr("COALESCE(post_count, 0) + ?", 1)).
		Error
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("AfterDelete hook triggered")
	var post Post
	if err := tx.Model(&Post{}).Where("id = ?", c.PostId).First(&post).Error; err != nil {
		return err
	}
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&post).UpdateColumn("comment_status", "无评论").Error
	}
	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// 把PostId等信息存到钩子自身的字段里，确保AfterDelete可用
	// 这里 c.PostId 就不会是空的，因为 BeforeDelete 是删除前调用的
	return nil
}
