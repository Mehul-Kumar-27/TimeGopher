package main

import (
	"fmt"
	task "timegopher/models"
	"flag"
)

func main() {
	var task task.Task
	task.Name = "Laundry"
	task.Description = "Wash the laundry"
	task.Completed = false

	fmt.Println(task.Name)

	flag.StringVar(&task.Name, "name", "", "Name of the task")
	flag.StringVar(&task.Description, "description", "", "Description of the task")
	flag.BoolVar(&task.Completed, "completed", false, "Whether or not the task is completed")

	flag.Parse()
	fmt.Println(task.Name)
}
