package user

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/the_one/code"
	"github.com/zhangshanwen/the_one/initialize/db"
	"github.com/zhangshanwen/the_one/initialize/service"
	"github.com/zhangshanwen/the_one/internal/param"
	"github.com/zhangshanwen/the_one/internal/response"
	"github.com/zhangshanwen/the_one/model"
	"github.com/zhangshanwen/the_one/tools"
)

func Login(c *service.Context) (resp service.Res) {
	p := param.Login{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{Mobile: p.Mobile}
	g := db.G
	if resp.Err = g.Where(&user).Preload("Wallet").First(&user).Error; resp.Err != nil {
		return
	}
	if !user.CheckPassword(p.Password) {
		resp.ResCode = code.ActPWdError
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &user); resp.Err != nil {
		return
	}
	var token string
	token, resp.Err = tools.CreateToken(user.Id)
	if resp.Err != nil {
		return
	}
	if user.Wallet == nil {
		user.Wallet = &model.Wallet{Uid: user.Id}
		if resp.Err = g.Create(&user.Wallet).Error; resp.Err != nil {
			return
		}
	}
	r.Balance = user.Wallet.Balance
	r.Authorization = token
	resp.Data = r
	return
}
