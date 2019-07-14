package main

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server.
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server.
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	// Sender data.
	from := "youremail@gmail.com"
	password := "yourSecretPassword"

	// Receiver email address.
	to := []string{"firstemail@gmail.com", "secondemail@gmail.com"}

	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	// Message.
	message := []byte("This is a really unimaginative message, I know.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent!")
}
