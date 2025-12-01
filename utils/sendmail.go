package utils

import (
    "net/smtp"
)

func SendEmail(to, subject, body string) error {
    from := "your@email.com"
    password := "YOUR_APP_PASSWORD"

    host := "smtp.gmail.com"
    port := "587"

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    auth := smtp.PlainAuth("", from, password, host)

    return smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
}
