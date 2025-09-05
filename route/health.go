package route

import "github.com/gin-gonic/gin"

func registerHealthCheck(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.Writer.WriteString("ok")
	})
}
