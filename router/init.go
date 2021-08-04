package router

import (
	"fmt"

	"github.com/zhangshanwen/the_one/initialize/app"
	"github.com/zhangshanwen/the_one/initialize/conf"
	"github.com/zhangshanwen/the_one/router/api"
)

func InitRouter() {
	api.RegisterApiV1Router()
	run()
}

func run() {
	_ = app.R.Run(fmt.Sprintf(":%s", conf.C.Port))

}
