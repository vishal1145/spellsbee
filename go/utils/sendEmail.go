
package utils

import (
	"gopkg.in/gomail.v2"
	"os"
)

 
func SendEmail(to, subject, body string) error {
	// Create new message
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Create dialer with Gmail settings
	d := gomail.NewDialer("smtp.gmail.com", 587, 
		os.Getenv("EMAIL_USER"), 
		os.Getenv("EMAIL_PASS"))

	// Send the email
	return d.DialAndSend(m)
}