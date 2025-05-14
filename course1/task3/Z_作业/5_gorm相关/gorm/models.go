package main

import "time"

type User struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name       string    `json:"name" gorm:"not null;column:name"`
	Age        int       `json:"age" gorm:"not null;column:age"`
	LastUpdate time.Time `json:"last_update" gorm:"not null;column:last_update"`

	Posts    []Post    `json:"posts" gorm:"foreignKey:UserId;references:Id"`
	Comments []Comment `json:"comments" gorm:"foreignKey:UserId;references:Id"`
}

type Post struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserId     int64     `json:"user_id" gorm:"not null;column:user_id;index"`
	Title      string    `json:"title" gorm:"not null;column:title"`
	Content    string    `json:"content" gorm:"not null;column:content"`
	LastUpdate time.Time `json:"last_update" gorm:"not null;column:last_update"`

	Comments []Comment `json:"comments" gorm:"foreignKey:PostId;references:Id"`
}

type Comment struct {
	Id         int64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserId     int64     `json:"user_id" gorm:"not null;column:user_id;index"`
	PostId     int64     `json:"post_id" gorm:"not null;column:post_id;index"`
	Content    string    `json:"content" gorm:"not null;column:content"`
	LastUpdate time.Time `json:"last_update" gorm:"not null;column:last_update"`
}
