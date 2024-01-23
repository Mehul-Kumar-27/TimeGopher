package commands

import (
	"time"
	task "timegopher/models"
)

func AddTask(name string, description string, completed bool, startTime time.Time) task.Task {
	var task task.Task
	task.Name = name
	task.Description = description
	task.Completed = completed

	return task
}