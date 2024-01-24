package commands

import (
	"fmt"
	"time"
	task "timegopher/models"
)

func AddTask(name string, description string, completed bool, startTime time.Time) task.Task {
	var task task.Task
	fmt.Println("Please enter the name of the task:")
	fmt.Scanln(&task.Name)
	fmt.Println("Please enter the description of the task:")
	fmt.Scanln(&task.Description)


	return task
}