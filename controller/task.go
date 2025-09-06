package controller

import (
	"gin-demo-framework/data"
	"gin-demo-framework/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: service.NewTaskService(data.GetDB()),
	}
}

func (t *TaskController) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params service.TaskListParams
		if err := ctx.Bind(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}

		tasks, total, err := t.taskService.List(ctx, &params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":   200,
			"data":   tasks,
			"total":  total,
			"params": params,
		})
	}
}

func (t *TaskController) Detail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, exist := ctx.Params.Get("id")
		if !exist {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "id is required",
			})
			return
		}
		taskID, _ := strconv.Atoi(id)
		taskDetail, err := t.taskService.Detail(ctx, taskID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "internal server error",
				"err":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": taskDetail,
		})
	}
}

func (t *TaskController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": 1,
		})
	}
}

func (t *TaskController) UpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": 1,
		})
	}
}

func (t *TaskController) DeleteTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": 1,
		})
	}
}
