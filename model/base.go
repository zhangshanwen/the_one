package model

type BaseModel struct {
	Id          int  `json:"id"             gorm:"primaryKey;autoIncrement"`
	CreatedTime int  `json:"created_time"   gorm:"autoCreateTime"` //autoUpdateTime:nano/milli 纳/毫秒
	UpdatedTime int  `json:"updated_time"   gorm:"autoUpdateTime"`
	IsDeleted   bool `json:"-"              gorm:"default:0"`
}
