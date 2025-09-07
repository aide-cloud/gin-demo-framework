package middler

import (
	"gin-demo-framework/controller/response"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("Recover middleware start")
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Recover middleware", "error", err)
				response.ErrorInternal(c, err)
			}
		}()
		c.Next()
		slog.Info("Recover middleware end")
	}
}
