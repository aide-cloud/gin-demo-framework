package route

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()

	// 注册路由
	registerHealthCheck(r)

	return r
}
