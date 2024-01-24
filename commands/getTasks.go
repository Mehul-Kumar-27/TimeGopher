package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	customLogger "timegopher/logger"
	t "timegopher/models"
)

func GetTasks() {
	customLogger := customLogger.CreateCustomLogger()

	// Read existing content from the file
	existingContent, err := os.ReadFile("tasks.json")
	if err != nil && !os.IsNotExist(err) {
		customLogger.Fatal(err)
	}

	// Unmarshal existing content
	var tasks []t.Task
	if len(existingContent) > 0 {
		err = json.Unmarshal(existingContent, &tasks)
		if err != nil {
			customLogger.Fatal(err)
		}
	}

	// Use tabwriter to format the output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Print header
	fmt.Fprintf(w, "%-30s\t%-50s\t%-10v\n", "NAME", "DESCRIPTION", "COMPLETED")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------------------------------------------")

	// Print tasks
	for _, task := range tasks {
		fmt.Fprintf(w, "%-30s\t%-50s\t%-10v\n", task.Name, task.Description, task.Completed)
	}

	// Flush the buffer
	w.Flush()
}
