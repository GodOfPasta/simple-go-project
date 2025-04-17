package main

import (
	"log"
	"simple-go-project/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/tasks", utils.AddTask)
	r.GET("/tasks", utils.GetTasks)
	r.GET("/tasks/:id", utils.GetTask)
	r.PUT("/tasks/:id", utils.UpdateTask)
	r.DELETE("/tasks/:id", utils.DeleteTask)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
