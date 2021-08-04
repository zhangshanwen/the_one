package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/zhangshanwen/the_one/initialize/conf"
	l "github.com/zhangshanwen/the_one/initialize/logger"
)

var R *redis.Client

func InitRedis() {
	l.Logger.Info("--------init_redis_client_start---------")
	redisConf := conf.C.DB.Redis
	host := redisConf.Host
	port := redisConf.Port
	if host == "" {
		host = "localhost"
	}
	if port <= 0 {
		port = 6379
	}
	R = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%v", host, port),
		Password:    redisConf.Password,
		DB:          redisConf.DataBase,
		ReadTimeout: time.Duration(redisConf.Timeout),
	})
	l.Logger.Info("--------init_redis_client_end---------")
}
