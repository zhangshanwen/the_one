package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/the_one/initialize/service"
)

func Handel(fun func(c *service.Context) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.Json(c, fun(&service.Context{Context: c}))
	}
}
