package main

import (
	"github.com/zhangshanwen/the_one/initialize"
	"github.com/zhangshanwen/the_one/router"
)

func main() {
	initialize.Initialize() // 注册服务
	router.InitRouter()     // 注册路由
}
