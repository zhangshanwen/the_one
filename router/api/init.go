package api

import (
	"github.com/zhangshanwen/the_one/initialize/app"
	"github.com/zhangshanwen/the_one/router/api/v1"
)

func RegisterApiV1Router() {
	api := app.R.Group("api")
	group := api.Group("v1")
	v1.InitVersion(group)
	v1.InitUserRouter(group)
}
