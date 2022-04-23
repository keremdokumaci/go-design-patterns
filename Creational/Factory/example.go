package main

import "fmt"

func main() {
	notification, err := NewNotification(Email, "email content", "email title")
	if err != nil {
		fmt.Errorf(err.Error())
	}

	smsNotification, err := NewNotification(Sms, "sms content", "sms title")
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(notification, "\n", smsNotification)
}
