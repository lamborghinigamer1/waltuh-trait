package debugging

import (
	"fmt"
	"time"
	"waltuh/colors"
)

type Logger struct {
	structName string
}

func NewLogger(structName string) *Logger {
	return &Logger{
		structName: structName,
	}
}

func (l *Logger) log(level string, message string) {
	formattedMessage := fmt.Sprintf("[%s - %s]: %s", level, getTime(), message)

	switch level {
	case "INFO":
		fmt.Printf("%s%s%s\n", colors.OkBlue, formattedMessage, colors.End)
	case "DEBUG":
		fmt.Printf("%s%s%s\n", colors.OkGreen, formattedMessage, colors.End)
	case "WARNING":
		fmt.Printf("%s%s%s\n", colors.Warning, formattedMessage, colors.End)
	case "FATAL":
		fmt.Printf("%s%s%s\n", colors.Header, formattedMessage, colors.End)
	case "ERROR":
		fmt.Printf("%s%s%s\n", colors.Fail, formattedMessage, colors.End)
	}
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Fatal(message string) {
	l.log("FATAL", message)
}

func (l *Logger) Warning(message string) {
	l.log("WARNING", message)
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func getTime() string {
	localTime := time.Now().Local()
	return localTime.Format("2006-01-02 15:04:05")
}
