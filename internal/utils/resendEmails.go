package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

func SendTokenWithResend(email, token string) {
    apiKey := os.Getenv("RESEND_API_KEY")

    client := resend.NewClient(apiKey)

	var message = fmt.Sprintf("<p>Your are trying to login to ai chat <strong>%s</strong>!</p>", token)

    params := &resend.SendEmailRequest{
        From:    "onboarding@resend.dev",
        To:      []string{email},
        Subject: "Welcome to the world of AI",
        Html:    message,
    }

    _, err := client.Emails.Send(params)

	if err != nil {
		log.Println(err)
	}else{
		log.Println("Email sent")
	}
}