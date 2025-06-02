package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendTokenEmail(email, token string) {

	from := "tylerjusfly1@gmail.com"
	password := os.Getenv("GOOGLE_APP_PASSWORD")

	to := []string{"tylerjusfly1@gmail.com"}

	// Subject and body of the email
	subject := "One-time login token"
	body := "This is a test email sent from Go using Gmail."

	message := []byte(fmt.Sprintf("To: %s\r\n", to[0]))
	message = append(message, []byte(fmt.Sprintf("From: %s\r\n", from))...)
	message = append(message, []byte(fmt.Sprintf("Subject: %s\r\n", subject))...)
	message = append(message, []byte("\r\n")...)
	message = append(message, []byte(body)...)

	// Configure SMTP server and port
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Send the email
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}

}
