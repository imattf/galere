// Fun with SMTP and Email

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/imattf/galere/models"
	"github.com/joho/godotenv"
)

func main() {

	// Using .env file...
	fmt.Println("Hello .env stuff...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	// Email setup testing...
	fmt.Println("Hello email stuff...")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})

	err = es.ForgotPassword("bob@aol.com", "https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}

	fmt.Println("...Message Sent")
}
