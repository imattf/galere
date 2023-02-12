// Fun with SMTP and Email

package main

import (
	"fmt"

	"github.com/imattf/galere/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "b5316f30395b9d"
	password = "348e7e664ef1a6"
)

func main() {
	fmt.Println("Hello email stuff...")

	email := models.Email{
		From:      "test@faulkners.io",
		To:        "bob@aol.com",
		Subject:   "This is a test email",
		Plaintext: "This the body of the email",
		HTML:      `<h1>Hi Bob!</h1><p>This is email</p><p>Please enjoy</p>`,
	}
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.Send(email)
	if err != nil {
		panic(err)
	}

	// from := "test@faulkners.io"
	// to := "bob@aol.com"
	// subject := "This is a test email"
	// plaintext := "This the body of the email"
	// html := `<h1>Hi Bob!</h1><p>This is email</p><p>Please enjoy</p>`

	// msg := mail.NewMessage()
	// msg.SetHeader("To", to)
	// msg.SetHeader("From", from)
	// msg.SetHeader("Subject", subject)
	// msg.SetBody("text/plain", plaintext)
	// msg.AddAlternative("text/html", html)
	// // msg.WriteTo(os.Stdout)

	// dialer := mail.NewDialer(host, port, username, password)

	// Dial and SendCloser method...
	// sender, err := dialer.Dial()
	// if err != nil {
	// 	// TODO: Handle the error correctly
	// 	panic(err)
	// }
	// defer sender.Close()
	// err = sender.Send(from, []string{to}, msg)
	// if err != nil {
	// 	// TODO Handle the error correctly
	// 	panic(err)
	// }
	// fmt.Println("...Message Sent")

	// DialandSend method...
	// err := dialer.DialAndSend(msg)
	// if err != nil {
	// 	// TODO: Handle the error correctly
	// 	panic(err)
	// }

	fmt.Println("...Message Sent")
}