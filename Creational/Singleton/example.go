package main

import "fmt"

func main() {
	logger := NewLoggerInstance()

	logger.SetLogLevel(1)
	logger.Log("Log level 1 message")

	logger.SetLogLevel(2)
	logger.Log("Log level 2 message")

	for i := 0; i < 10; i++ {
		go func(level int) {
			logger.SetLogLevel(level)
			logger.Log("Log from different routine.")
		}(i)
	}
	fmt.Scanln()

}
