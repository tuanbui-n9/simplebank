package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddr   = "smtp.gmail.com"
	smtpServerAddr = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddr     string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddr, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddr:     fromEmailAddr,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddr)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.HTML = []byte(content)

	for _, file := range attachFiles {
		if _, err := e.AttachFile(file); err != nil {
			return fmt.Errorf("failed to attach file: %w", err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddr, sender.fromEmailPassword, smtpAuthAddr)
	return e.Send(smtpServerAddr, smtpAuth)
}
