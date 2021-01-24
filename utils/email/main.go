package Email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/tamnk74/todolist-mysql-go/config"
)

const from = "super_admin@gmail.com"

func Send(to []string) {
	// Authentication.
	auth := smtp.PlainAuth("", config.MAIL.USERNAME, config.MAIL.PASSWORD, config.MAIL.HOST)

	t, _ := template.ParseFiles("templates/email/welcome.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name string
	}{
		Name: "Tom NK",
	})

	// Sending email.
	err := smtp.SendMail(config.MAIL.HOST+":"+config.MAIL.PORT, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("An email has been sent from ", from, " to ", to)
}
