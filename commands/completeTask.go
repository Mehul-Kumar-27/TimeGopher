package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"
	"time"
	customLogger "timegopher/logger"
	t "timegopher/models"

	"github.com/eiannone/keyboard"
)

func SelectAndCompleteTask() {
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
	fmt.Fprintln(w, "Index\tName\tDescription\tCompleted")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------------------------------------------")

	// Print tasks with index
	for i, task := range tasks {
		fmt.Fprintf(w, "%-6d\t%-30s\t%-50s\t%-10v\n", i+1, task.Name, task.Description, task.Completed)
	}

	// Flush the buffer
	w.Flush()

	err = keyboard.Open()
	if err != nil {
		customLogger.Fatal(err)
	}
	defer func() {
		// Close the keyboard when done
		_ = keyboard.Close()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		// Listen for Ctrl-C and close the keyboard
		<-ch
		customLogger.Fatalln("Received Ctrl-C, exiting.")
	}()

	// Wait for user input
	fmt.Println("\nUse arrow keys to select a task, then press Enter to mark it as completed.")
	index := 0
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			customLogger.Fatal(err)
		}

		if key == keyboard.KeyEnter {
			break
		}

		if key == keyboard.KeyArrowUp && index > 0 {
			index--
		}

		if key == keyboard.KeyArrowDown && index < len(tasks)-1 {
			index++
		}

		// Clear the console
		fmt.Print("\033[H\033[2J")

		// Use tabwriter to format the output with the selected task highlighted
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		// Print header
		fmt.Fprintln(w, "Index\tName\tDescription\tCompleted")
		fmt.Fprintln(w, "--------------------------------------------------------------------------------------------------------------------")

		// Print tasks with index and highlight the selected task
		for i, task := range tasks {
			if i == index {
				fmt.Fprintf(w, "\033[1m%-6d\t%-30s\t%-50s\t%-10v\033[0m\n", i+1, task.Name, task.Description, task.Completed)
			} else {
				fmt.Fprintf(w, "%-6d\t%-30s\t%-50s\t%-10v\n", i+1, task.Name, task.Description, task.Completed)
			}
		}

		// Flush the buffer
		w.Flush()

		// Sleep for a short duration to avoid capturing multiple key presses
		time.Sleep(50 * time.Millisecond)
	}

	// Mark the selected task as completed
	if index >= 0 && index < len(tasks) {
		tasks[index].Completed = true

		// Save the updated tasks to the file
		updatedContent, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			customLogger.Fatal(err)
		}

		err = os.WriteFile("tasks.json", updatedContent, 0644)
		if err != nil {
			customLogger.Fatal(err)
		}

		fmt.Printf("Task '%s' marked as completed.\n", tasks[index].Name)
	} else {
		fmt.Println("Invalid task selection.")
	}
}
