package app

import (
	"github.com/gin-gonic/gin"
)

var R = &gin.Engine{}

func init() {
	gin.ForceConsoleColor()
	R = gin.Default()
}
