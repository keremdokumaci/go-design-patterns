package main

import "fmt"

func main() {
	builder := NewNotificationBuilder()
	builder.setTitle("title")
	builder.setContent("content")
	builder.setRetryCount(2)
	builder.setSubtitle("subtitle")

	notification, err := builder.Build()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(notification)
}
