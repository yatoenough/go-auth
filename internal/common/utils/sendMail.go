package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(receiver, activationCode string) {
	//get env variables
	senderMail := os.Getenv("MAILER_SENDER_MAIL")
	appPassword := os.Getenv("MAILER_APP_PASSWORD")
	smtpHost := os.Getenv("MAILER_SMTP_HOST")
	smtpPort := os.Getenv("MAILER_SMTP_PORT")

	//connect to SMTP
	auth := smtp.PlainAuth("", senderMail, appPassword, smtpHost)
	to := []string{receiver}

	activationLink := os.Getenv("API_HOST") + "/api/v1/users/activate/" + activationCode

	//define message context
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Activate your account\r\n"+
		"\r\n"+
		"Activate your account by following this link: %s\r\n", receiver, activationLink))

	//send mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderMail, to, msg)
	if err != nil {
		fmt.Println(err)
	}
}
