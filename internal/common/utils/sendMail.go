package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(receiver, activationCode string) error {
	apiHost := os.Getenv("API_HOST")
	senderMail := os.Getenv("SENDER_MAIL")
	appPassword := os.Getenv("APP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", senderMail, appPassword, smtpHost)
	to := []string{receiver}

	activationLink := apiHost + "/api/v1/users/activate/" + activationCode

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Activate your account\r\n"+
		"\r\n"+
		"Activate your account by following this link: %s\r\n", receiver, activationLink))

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderMail, to, msg)
	if err != nil {
		return err
	}

	return nil
}
