package main

import "errors"

type NotificationType string

const (
	Email NotificationType = "email"
	Sms                    = "sms"
)

func NewNotification(notiType NotificationType, content string, title string) (INotification, error) {
	if notiType == Email {
		return CreateEmailNotification(content, title), nil
	}
	if notiType == Sms {
		return CreateSmsNotification(content, title), nil
	}

	return nil, errors.New("couldn't find notification")
}
