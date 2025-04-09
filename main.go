package main

import (
	"log"
	"simple-go-project/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", utils.Test)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
