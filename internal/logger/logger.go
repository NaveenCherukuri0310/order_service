package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	// Open or create log file
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create logger with flags
	Log = log.New(f, "", log.LstdFlags|log.Lshortfile)
}
