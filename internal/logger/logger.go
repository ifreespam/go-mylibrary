package logger

// CustomLogger is a custom logger instance
type CustomLogger struct {
	label string
}

// Log function prints the string data
func (cl *CustomLogger) Log(str string) string {
	var logString = "[" + cl.label + "]: " + str
	return logString
}

// GetLogger is a factory to create the logger instance
func GetLogger(label string) *CustomLogger {
	return &CustomLogger{
		label: label,
	}
}
