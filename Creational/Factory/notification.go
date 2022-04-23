package main

type INotification interface {
	setTitle(title string)
	setContent(content string)
}

type Notification struct {
	title   string
	content string
}

func (n Notification) setTitle(title string) {
	n.title = title
}

func (n Notification) setContent(content string) {
	n.content = content
}
