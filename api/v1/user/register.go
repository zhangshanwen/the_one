package user

import (
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/zhangshanwen/the_one/code"
	"github.com/zhangshanwen/the_one/initialize/db"
	"github.com/zhangshanwen/the_one/initialize/service"
	"github.com/zhangshanwen/the_one/internal/param"
	"github.com/zhangshanwen/the_one/internal/response"
	"github.com/zhangshanwen/the_one/model"
)

func Register(c *service.Context) (resp service.Res) {
	p := param.Register{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{Mobile: p.Mobile}
	g := db.G
	g = g.Begin()
	defer func() {
		if resp.Err == nil {
			g.Commit()
		} else {
			g.Rollback()
		}
	}()
	if resp.Err = g.Where(&user).First(&user).Error; resp.Err != nil && resp.Err != gorm.ErrRecordNotFound {
		return
	}
	if user.Id > 0 {
		resp.Err = errors.New("mobile is existed")
		resp.ResCode = code.MobileIsExistEd
		return
	}
	if resp.Err = copier.Copy(&user, &p); resp.Err != nil {
		return
	}
	if resp.Err = user.SetPassword(p.Password); resp.Err != nil {
		return
	}
	if resp.Err = g.Save(&user).Error; resp.Err != nil {
		return
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &user); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
