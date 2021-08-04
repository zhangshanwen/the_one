package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/the_one/code"
	l "github.com/zhangshanwen/the_one/initialize/logger"
)

type Res struct {
	StatusCode int
	ResCode    int
	Data       interface{}
	Err        error
}

type res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func success(c *gin.Context, r Res) {
	c.JSON(http.StatusOK, res{
		Code: code.Success,
		Data: r.Data,
	})
}

func failed(c *gin.Context, r Res) {
	var msg string
	if gin.IsDebugging() {
		msg = r.Err.Error()
	}
	if r.ResCode == 0 {
		r.ResCode = code.BaseFailedCode
	}
	c.JSON(http.StatusBadRequest, res{
		Code: r.ResCode,
		Data: nil,
		Msg:  msg,
	})
}

func Json(c *gin.Context, r Res) {
	if r.StatusCode == 0 {
		if r.Err == nil {
			success(c, r)
		} else {
			failed(c, r)
		}
		return
	}
	var msg string
	if r.Err != nil {
		l.Logger.Error(r.Err)
		msg = r.Err.Error()
	}
	c.JSON(r.StatusCode, res{
		Code: r.ResCode,
		Data: r.Data,
		Msg:  msg,
	})
}
