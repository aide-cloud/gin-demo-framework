package middler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(apiList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		for _, api := range apiList {
			if strings.HasPrefix(path, api) {
				break
			}
		}
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "请输入密码"})
			return
		}
		if username != "admin" || password != "123456" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
			return
		}
		c.Next()
	}
}
