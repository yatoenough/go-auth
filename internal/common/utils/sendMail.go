package utils

import (
	"fmt"
	"go-auth/config"
	"net/smtp"
)

func SendMail(receiver, activationLink string) {
	auth := smtp.PlainAuth("", config.GetSenderMail(), config.GetAppPassword(), config.GetSmtpHost())
	to := []string{receiver}

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Activate your account\r\n"+
		"\r\n"+
		"Activate your account by following this link: %s\r\n", receiver, activationLink))

	err := smtp.SendMail(config.GetSmtpHost()+":"+config.GetSmtpPort(), auth, config.GetSenderMail(), to, msg)

	if err != nil {
		fmt.Println(err)
	}
}
