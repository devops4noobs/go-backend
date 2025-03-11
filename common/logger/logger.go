// common/logger/logger.go
package logger

import (
	"log"
	"os"
)

// Initialize logger with output to a file
func InitLogger() *log.Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open log file, using default stderr")
		return log.New(os.Stderr, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
