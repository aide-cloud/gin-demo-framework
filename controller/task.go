package controller

import (
	"gin-demo-framework/data/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (t *TaskController) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": []*model.Task{},
	})
}

func (t *TaskController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": 1,
		})
	}
}
