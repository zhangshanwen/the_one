package logger

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *log.Logger

func InitGinLogger() {
	gin.DefaultWriter = Logger.Out
	Logger.Info("......GIN日志初始化成功......")
}

func InitLog() {
	Logger = log.New()
	Logger.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	//
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "log/the_one.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     7, // days
	})
	Logger.SetOutput(mw)
	InitGinLogger()
}
