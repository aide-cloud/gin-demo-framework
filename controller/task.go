package controller

import (
	"gin-demo-framework/controller/request"
	"gin-demo-framework/controller/response"
	"gin-demo-framework/data"
	"gin-demo-framework/service"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: service.NewTaskService(data.GetDB(), data.GetRDB()),
	}
}

func (t *TaskController) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params service.TaskListParams
		if err := request.Bind(ctx, &params); err != nil {
			response.ErrorParams(ctx, err)
			return
		}

		taskListReply, err := t.taskService.List(ctx, &params)
		if err != nil {
			response.ErrorBusiness(ctx, err)
			return
		}
		response.Success(ctx, taskListReply)
	}
}

type TaskDetailParams struct {
	ID int `form:"id" uri:"id" json:"id"`
}

func (t *TaskController) Detail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// panic("test panic")
		var params TaskDetailParams
		if err := request.Bind(ctx, &params); err != nil {
			response.ErrorParams(ctx, err)
			return
		}

		taskDetail, err := t.taskService.Detail(ctx, params.ID)
		if err != nil {
			response.ErrorBusiness(ctx, err)
			return
		}
		response.Success(ctx, taskDetail)
	}
}

func (t *TaskController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response.Success(ctx, 1)
	}
}

func (t *TaskController) UpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response.Success(ctx, 1)
	}
}

func (t *TaskController) DeleteTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response.Success(ctx, 1)
	}
}
