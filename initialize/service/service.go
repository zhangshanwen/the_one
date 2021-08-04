package service

import (
	"github.com/gin-gonic/gin"

	l "github.com/zhangshanwen/the_one/initialize/logger"
	"github.com/zhangshanwen/the_one/model"
)

type Context struct {
	*gin.Context
	User model.User
}

func (c *Context) Rebind(obj interface{}) (err error) {
	if err = c.Bind(obj); err != nil {
		return
	}
	l.Logger.WithField("mod", "params").Info(obj)
	return
}
