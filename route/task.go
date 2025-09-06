package route

import (
	"gin-demo-framework/controller"

	"github.com/gin-gonic/gin"
)

func registerTaskRoutes(router *gin.RouterGroup) {
	taskController := controller.NewTaskController()
	router.GET("/tasks", taskController.List())
	router.GET("/task/:id", taskController.Detail())
	router.POST("/tasks", taskController.Create())
	router.PUT("/tasks/:id", taskController.UpdateTask())
	router.DELETE("/tasks/:id", taskController.DeleteTask())
}
