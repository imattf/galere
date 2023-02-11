// Fun with SMTP and Email

package main

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "b509d12c91911f"
	password = "90171c99a302fd"
)

func main() {
	fmt.Println("Hello email stuff...")

	from := "test@faulkners.io"
	to := "bob@aol.com"
	subject := "This is a test email"
	plaintext := "This the body of the email"
	html := `<h1>Hi Bob!</h1><p>This is email</p><p>Please enjoy</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	// msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)

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
	err := dialer.DialAndSend(msg)
	if err != nil {
		// TODO: Handle the error correctly
		panic(err)
	}

	fmt.Println("...Message Sent")
}
