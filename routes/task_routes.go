package routes

import (
	"crud-api/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
