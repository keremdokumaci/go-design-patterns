package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	logLevel int
}

func (l *Logger) Log(log string) {
	fmt.Println("Log Level : ", l.logLevel, " Log : ", log)
}

func (l *Logger) SetLogLevel(level int) {
	l.logLevel = level
}

var logger *Logger

var once sync.Once

func NewLoggerInstance() *Logger {
	once.Do(func() {
		fmt.Println("Creating logger instance ..")
		logger = &Logger{}
	})

	return logger
}
