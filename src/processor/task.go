package processor

import (
	"server/src/models"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	TasksQueue = make(chan string, 100)
	Tasks      sync.Map
)

func StartWork() {
	for id := range TasksQueue {
		go ProcessTask(id)
	}
}

func ProcessTask(id string) {
	TaskAny, ok := Tasks.Load(id)
	if !ok {
		return
	}

	task := TaskAny.(*models.Task)

	logrus.Info("processing...")
	time.Sleep(3 * time.Minute)

	res := "task " + id + " is done"
	task.Result = res
	task.Status = models.StatusDone
	Tasks.Store(id, task)
}

func AddTask(task *models.Task) {
	Tasks.Store(task.Id, task)
	TasksQueue <- task.Id
}

func GetTask(id string) (*models.Task, bool) {
	TaskAny, ok := Tasks.Load(id)
	if !ok {
		return nil, false
	}
	return TaskAny.(*models.Task), true
}
