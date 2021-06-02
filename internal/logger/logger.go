package logger

import "fmt"

// CustomLogger is a custom logger instance
type CustomLogger struct {
	label string
}

// Log function prints the string data
func (cl *CustomLogger) Log(str string) {
	var logString = "[" + cl.label + "]: " + str
	fmt.Println(logString)
}

// GetLogger is a factory to create the logger instance
func GetLogger(label string) *CustomLogger {
	return &CustomLogger{
		label: label,
	}
}
