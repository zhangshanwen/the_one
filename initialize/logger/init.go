package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *log.Logger

func InitGinLogger() {
	gin.DefaultWriter = Logger.Writer()
	Logger.Info("......GIN日志初始化成功......")
}

func InitLog() {
	Logger = log.New()
	Logger.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	//
	Logger.SetOutput(os.Stdout)
	// 输出到文件
	Logger.SetOutput(&lumberjack.Logger{
		Filename:   "log/the_one.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})
	InitGinLogger()
}
