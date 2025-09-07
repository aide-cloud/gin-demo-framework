package middler

import (
	"gin-demo-framework/data/model"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Logger(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("Logger middleware start")
		start := time.Now()
		c.Next()
		slog.Info("Logger middleware end", "status", c.Writer.Status(), "latency", time.Since(start), "path", c.Request.URL.Path, "method", c.Request.Method, "clientIP", c.ClientIP())
		db.Create(&model.Log{
			Status:   c.Writer.Status(),
			Latency:  time.Since(start),
			Path:     c.Request.URL.Path,
			Method:   c.Request.Method,
			ClientIP: c.ClientIP(),
		})
	}
}
