package model

import (
	_ "gorm.io/gorm"
	"time"
)

type Model struct {
	ID uint `gorm:"column:id; PRIMARY_KEY"`
}

type New struct {
	Model
	Symbol    string    `gorm:"column:symbol"`
	Token     string    `gorm:"column:token"`
	Price     float64   `gorm:"column:price"`
	Pair      string    `gorm:"column:pair"`
	Tweet     int       `gorm:"column:tweet"`
	Liquidity float64   `gorm:"column:liquidity"`
	Age       string    `gorm:"column:age"`
	PushTime  time.Time `gorm:"column:push_time"`
	Times     int64     `gorm:"column:times"`
}

type Hot struct {
	Model
	Symbol    string    `gorm:"column:symbol"`
	Token     string    `gorm:"column:token"`
	Price     float64   `gorm:"column:price"`
	Pair      string    `gorm:"column:pair"`
	Vol       float64   `gorm:"column:vol"`
	Tx        int       `gorm:"column:tx"`
	Trader    int       `gorm:"column:trader"`
	Liquidity float64   `gorm:"column:liquidity"`
	Age       string    `gorm:"column:age"`
	Tweet     int       `gorm:"column:tweet"`
	PushTime  time.Time `gorm:"column:push_time"`
	Times     int64     `gorm:"column:times"`
	Type      int64     `gorm:"column:type"`
	Level     string    `gorm:"column:level"`
}
