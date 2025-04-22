package handlers

import (
	"server/src/models"
	"server/src/processor"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTaskStatus(c *gin.Context) {
	id := c.Param("id")

	task, ok := processor.GetTask(id)
	if !ok {
		c.JSON(404, gin.H{
			"error": "task not found",
		})
	}

	c.JSON(200, task)
}

func PostTask(c *gin.Context) {
	id := uuid.New().String()
	task := &models.Task{
		Id:     id,
		Status: models.StatuseProcessing,
	}
	processor.AddTask(task)

	c.JSON(200, task)
}
