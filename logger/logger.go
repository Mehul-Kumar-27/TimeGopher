package logger

import (
	"log"
	"os"
)

func CreateCustomLogger() *log.Logger {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger	
}
