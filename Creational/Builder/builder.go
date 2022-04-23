package main

type NotificationBuilder struct {
	Title      string
	Content    string
	Subtitle   string
	RetryCount int
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (bldr *NotificationBuilder) setTitle(title string) {
	bldr.Title = title
}

func (bldr *NotificationBuilder) setSubtitle(subTitle string) {
	bldr.Subtitle = subTitle
}

func (bldr *NotificationBuilder) setContent(content string) {
	bldr.Content = content
}

func (bldr *NotificationBuilder) setRetryCount(count int) {
	bldr.RetryCount = count
}

func (bldr *NotificationBuilder) Build() (*Notification, error) {
	ntf := Notification{}
	ntf.title = bldr.Title
	ntf.content = bldr.Content
	ntf.subtitle = bldr.Subtitle
	ntf.retryCount = bldr.RetryCount
	return &ntf, nil
}
