package logging

import "fmt"

type ConsoleLogger struct {
}

func NewConsoleLogger() Logger {
	return ConsoleLogger{}
}

func (l ConsoleLogger) Log(le LogEvent) {
	fmt.Println(le)
}
