package model

import "time"

type Model struct {
	ID uint `gorm:"column:id; PRIMARY_KEY"`
}

type TweetAnalysePushLog struct {
	Model
	Keyword  string    `gorm:"column:keyword"`
	ListId   string    `gorm:"column:list_id"`
	ListName string    `gorm:"column:list_name"`
	Times    int64     `gorm:"column:times"`
	Tweet    int64     `gorm:"column:tweet"`
	PushTime time.Time `gorm:"column:push_time"`
}
