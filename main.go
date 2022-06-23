package main

import (
	"flip/aws-sdk-go-poc/models"
	v1 "flip/aws-sdk-go-poc/v1"
	v2 "flip/aws-sdk-go-poc/v2"
	"fmt"
)

const (
	contentTitle = `Testing Email`
	contentText  = `Content of Testing Email`
	contentHTML  = `<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with 
	<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the 
	<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>`
)

func main() {

	request := models.RequestSendEmail{
		To:      []string{"your_email@gmail.com"},
		Cc:      []string{},
		Bcc:     []string{},
		From:    "your_email@gmail.com",
		Subject: contentTitle,
		Body:    contentHTML,
		Type:    models.HTMLTypeEmail,
	}

	if err := v1.SendEmail(request); err != nil {
		fmt.Printf("Sending email error %s", err.Error())
	}

	if err := v2.SendEmail(request); err != nil {
		fmt.Printf("Sending email error %s", err.Error())
	}
}
