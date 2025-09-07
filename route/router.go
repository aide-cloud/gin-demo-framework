package route

import (
	"gin-demo-framework/config"
	"gin-demo-framework/data"
	"gin-demo-framework/route/middler"

	"github.com/gin-gonic/gin"
)

func New(c *config.ServerConfig) *gin.Engine {
	r := gin.New()
	// 注册路由
	registerHealthCheck(r)

	gl := r.Group(c.Prefix)

	// 注册全局中间件
	gl.Use(middler.Recover())
	gl.Use(middler.Logger(data.GetDB()))
	gl.Use(middler.Cors())

	api := gl.Group("api")
	api.Use(middler.BasicAuth([]string{"/api/tasks"}))
	// 注册任务路由
	registerTaskRoutes(api)

	return r
}
