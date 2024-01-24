package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	customLogger "timegopher/logger"
	t "timegopher/models"
)

func AddTask() t.Task {
	var task t.Task
	fmt.Print("Please enter the name of the task: ")
	fmt.Scanln(&task.Name)

	fmt.Print("Please enter the description of the task: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		task.Description = scanner.Text()
	}

	task.Completed = false

	saveTaskToFile(task)

	return task
}
func saveTaskToFile(task t.Task) {
	customLogger := customLogger.CreateCustomLogger()

	// Open the file in read-write mode, create it if it doesn't exist, and set permissions
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		customLogger.Fatal(err)
	}
	defer file.Close()

	var tasks []t.Task

	// Get file info to check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		customLogger.Fatal(err)
	}

	// If the file is not empty, decode its content
	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&tasks)
		if err != nil && err != io.EOF {
			customLogger.Fatal(err)
		}
	}

	// Append the new task
	tasks = append(tasks, task)

	// Reset the file cursor to the beginning
	file.Seek(0, 0)

	// Truncate the file to remove any existing content
	file.Truncate(0)

	// Encode and write the updated tasks to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		customLogger.Fatal(err)
	}

	customLogger.Println("Task saved to file 😊")
}
