package logging

import "fmt"

type ConsoleLogger struct {
}

func NewConsoleLogger() Logger {
	return FileLogger{}
}

func (l FileLogger) Log(le LogEvent) {
	fmt.Println(le)
}
