package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var cmd string
	flag.StringVar(&cmd, "command", "", "command to execute")

	flag.Parse()

	// Check if the command is provided
	if cmd == "" {
		fmt.Println("Please provide a command using the -command flag.")
		os.Exit(1)
	}

	fmt.Println("Command:", cmd)
}
