package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/the_one/api/v1/user"
	"github.com/zhangshanwen/the_one/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("user")
	v := middleware.Handel
	jwt := middleware.JwtHandel
	{
		r.POST("", v(user.Register))   // 创建用户
		r.POST("login", v(user.Login)) // 登录用户
		r.GET("", jwt(user.Detail))    // 获取用户信息
	}
}
