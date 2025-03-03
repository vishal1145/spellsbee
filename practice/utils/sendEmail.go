package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	emailUser := "muskan.shukla@algofolks.com"
	emailPass := "yazf fyzy mgno fbdb"

	if emailUser == "" || emailPass == "" {
		return fmt.Errorf("email credentials not set")
	}

	auth := smtp.PlainAuth("", emailUser, emailPass, smtpHost)

	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailUser, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
