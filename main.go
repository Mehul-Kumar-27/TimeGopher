package main

import (
	"flag"
	gopherCommands "timegopher/commands"
	customLogger "timegopher/logger"
)

func main() {
	var cmd string
	customLogger := customLogger.CreateCustomLogger()
	flag.StringVar(&cmd, "command", "", "command to execute")

	flag.Parse()

	switch cmd {
	case "add":
		gopherCommands.AddTask()
	case "list":
		gopherCommands.GetTasks()
	case "complete":
		gopherCommands.SelectAndCompleteTask()
	default:
		customLogger.Fatalln("Please provide a valid command")
	}
}
