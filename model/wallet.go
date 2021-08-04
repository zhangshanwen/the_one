package model

type Wallet struct {
	BaseModel
	Uid     int   `json:"uid"     gorm:"index"`
	Balance int64 `json:"balance"`
}
