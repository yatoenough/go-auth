package service

import (
	"fmt"
	"net/smtp"
)

type MailService interface {
	SendActivationMail(receiver, activationCode string) error
}

type mailServiceImpl struct {
	apiHost     string
	senderMail  string
	appPassword string
	smtpHost    string
	smtpPort    string
}

func NewMailService(apiHost, senderMail, appPassword, smtpHost, smtpPort string) MailService {
	return &mailServiceImpl{
		apiHost:     apiHost,
		senderMail:  senderMail,
		appPassword: appPassword,
		smtpHost:    smtpHost,
		smtpPort:    smtpPort,
	}
}

func (ms *mailServiceImpl) SendActivationMail(receiver, activationCode string) error {

	auth := smtp.PlainAuth("", ms.senderMail, ms.appPassword, ms.smtpHost)
	to := []string{receiver}

	activationLink := ms.apiHost + "/api/v1/users/activate/" + activationCode

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Activate your account\r\n"+
		"\r\n"+
		"Activate your account by following this link: %s\r\n", receiver, activationLink))

	err := smtp.SendMail(ms.smtpHost+":"+ms.smtpPort, auth, ms.senderMail, to, msg)
	if err != nil {
		return err
	}

	return nil
}
