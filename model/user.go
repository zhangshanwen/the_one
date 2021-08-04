package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Mobile   string  `json:"mobile"`
	NickName string  `json:"nick_name"`
	Password string  `json:"-"`
	Email    string  `json:"email"`
	Avatar   string  `json:"avatar"`
	Wallet   *Wallet `gorm:"foreignkey:Uid;rerences:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
