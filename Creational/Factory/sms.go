package main

import "fmt"

type SmsNotification struct {
	Notification
}

func (e SmsNotification) String() string {
	return fmt.Sprintf("Sms notification : %s", e.content)
}

func CreateSmsNotification(content string, title string) INotification {
	return SmsNotification{
		Notification: Notification{
			title:   title,
			content: content,
		},
	}
}
