package mail

import (
	"fmt"
	"net/smtp"
)

type nodeMailer struct {
	host     string
	port     string
	username string
	password string
}

func NewNodeMailer(host, port, username, password string) *nodeMailer {
	return &nodeMailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}
func (nm *nodeMailer) SendVerificationEmail(to string, token string) error {
	auth := smtp.PlainAuth("", nm.username, nm.password, nm.host)

	msg := []byte(fmt.Sprintf("Subject: Email Verification\n\nClick here to verify: https://yourdomain.com/verify?token=%s", token))

	addr := nm.host + ":" + nm.port

	err := smtp.SendMail(addr, auth, nm.username, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
