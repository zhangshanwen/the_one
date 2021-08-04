package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/the_one/code"
	"github.com/zhangshanwen/the_one/initialize/db"
	"github.com/zhangshanwen/the_one/initialize/service"
	"github.com/zhangshanwen/the_one/internal/header"
	"github.com/zhangshanwen/the_one/tools"
)

func Handel(fun func(c *service.Context) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.Json(c, fun(&service.Context{Context: c}))
	}
}
func verifyJwt(c *gin.Context) (res service.Res, claims *tools.Claims) {
	h := header.Authorization{}
	if res.Err = c.ShouldBindHeader(&h); res.Err != nil {
		res.StatusCode = http.StatusUnauthorized
		res.ResCode = code.AuthFailed
		return
	}
	claims, res.Err = tools.VerifyToken(h.Authorization)
	if res.Err != nil {
		return
	}
	return
}
func JwtHandel(fun func(*service.Context) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, claims := verifyJwt(c)
		if res.Err != nil {
			res.StatusCode = http.StatusUnauthorized
			res.ResCode = code.AuthFailed
			service.Json(c, res)
			return
		}
		sC := &service.Context{Context: c}
		if res.Err = db.G.First(&sC.User, claims.Payload.Uid).Error; res.Err != nil {
			res.StatusCode = http.StatusUnauthorized
			res.ResCode = code.AuthFailed
			service.Json(c, res)
			return
		}
		service.Json(c, fun(sC))
	}
}
