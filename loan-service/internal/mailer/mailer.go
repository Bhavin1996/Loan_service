package mailer

import (
	"crypto/tls"
	"net/smtp"
)

var smtpServer string
var smtpPort string
var smtpUser string
var smtpPassword string

func InitMailer(server, port, user, password string) {
	smtpServer = server
	smtpPort = port
	smtpUser = user
	smtpPassword = password
}

func SendEmail(toEmail, subject, body string) error {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	conn, err := tls.Dial("tcp", smtpServer+":"+smtpPort, tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		return err
	}

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(smtpUser); err != nil {
		return err
	}

	if err = client.Rcpt(toEmail); err != nil {
		return err
	}

	wc, err := client.Data()
	if err != nil {
		return err
	}

	_, err = wc.Write(msg)
	if err != nil {
		return err
	}

	err = wc.Close()
	if err != nil {
		return err
	}

	return client.Quit()
}
