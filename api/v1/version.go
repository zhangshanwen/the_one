package v1

import (
	"github.com/zhangshanwen/the_one/initialize/service"
)

var version string

func Version(c *service.Context) (resp service.Res) {
	resp.Data = version
	return
}
