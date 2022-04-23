package main

import "fmt"

type EmailNotification struct {
	Notification
}

func (e EmailNotification) String() string {
	return fmt.Sprintf("Email notification : %s", e.content)
}

func CreateEmailNotification(content string, title string) INotification {
	return EmailNotification{
		Notification: Notification{
			title:   title,
			content: content,
		},
	}
}
