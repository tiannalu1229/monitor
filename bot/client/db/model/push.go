package model

type PushLog struct {
	Model
	Token    string  `gorm:"column:token"`
	PushTime int64   `gorm:"column:time"`
	Times    int64   `gorm:"column:times"`
	Vol      float64 `gorm:"column:vol"`
	Tx       int64   `gorm:"column:tx"`
	Trader   int64   `gorm:"column:trader"`
	Tweet    int64   `gorm:"column:tweet"`
	PushType int64   `gorm:"column:push_type"`
}
