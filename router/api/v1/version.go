package v1

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/zhangshanwen/the_one/api/v1"
	"github.com/zhangshanwen/the_one/middleware"
)

func InitVersion(Router *gin.RouterGroup) {
	r := Router.Group("version")
	v := middleware.Handel
	{
		r.GET("", v(v1.Version)) // 获取当前版本
	}
}
