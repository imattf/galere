// Fun with SMTP and Email

package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
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
	msg.WriteTo(os.Stdout)
}
