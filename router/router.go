package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTask)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.NewTask)

	router.Run("localhost:8080")
}
