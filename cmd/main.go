package main

import (
	"server/src/handlers"
	"server/src/processor"

	"github.com/gin-gonic/gin"
)

func main() {
	go processor.StartWork()
	r := gin.Default()
	r.POST("/CreateTask", handlers.PostTask)
	r.GET("/GetTaskStatus/:id", handlers.GetTaskStatus)
	r.Run(":8090")
}
