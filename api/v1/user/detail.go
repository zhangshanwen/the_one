package user

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/the_one/initialize/service"
	"github.com/zhangshanwen/the_one/internal/response"
)

func Detail(c *service.Context) (resp service.Res) {
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &c.User); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
