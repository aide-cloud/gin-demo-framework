package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func ErrorParams(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  fmt.Sprintf("参数错误: %s", err),
	})
}

func ErrorBusiness(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  fmt.Sprintf("业务错误: %s", err),
	})
}

func ErrorInternal(ctx *gin.Context, err any) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  fmt.Sprintf("系统错误: %s", err),
	})
}
