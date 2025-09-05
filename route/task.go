package route

import (
	"gin-demo-framework/controller"

	"github.com/gin-gonic/gin"
)

func registerTaskRoutes(router *gin.RouterGroup) {
	taskController := controller.NewTaskController()
	router.GET("/tasks", taskController.List)
	router.POST("/tasks", taskController.Create())
	// router.PUT("/tasks/:id", updateTask)
	// router.DELETE("/tasks/:id", deleteTask)
}
