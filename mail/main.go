package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

const (
	mailUser = "tom@xxx.com"
	mailPWD  = "xxxxxxx"
	mailHost = "mail.xxx.com"
)

// UnencryptedAuth use unSSL to link mail server
type UnencryptedAuth struct {
	smtp.Auth
}

// Start overwrite
func (a UnencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	server.TLS = true
	return a.Auth.Start(server)
}

//Request mail request handler
type Request struct {
	auth    UnencryptedAuth
	to      []string
	subject string
}

// NewRequest Request constructer
func NewRequest(auth UnencryptedAuth, to []string, subject string) *Request {
	return &Request{
		auth:    auth,
		to:      to,
		subject: subject,
	}
}

// SendEmail send email
func (r *Request) SendEmail(body string) (bool, error) {
	msg := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s",
		strings.Join(r.to, ";"), mailUser, r.subject, body)

	addr := mailHost + ":25"

	if err := SendMail(addr, r.auth, mailUser, r.to, []byte(msg)); err != nil {
		return false, err
	}

	return true, nil
}

// SendMail for function SendEmail
func SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: mailHost, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}
	if a != nil {
		if ok, _ := c.Extension("AUTH"); !ok {
			return errors.New("smtp: server doesn't support AUTH")
		}
		if err = c.Auth(a); err != nil {
			return err
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// Mail for outter
func Mail(mailTo []string, title string, body string) (bool, error) {
	auth := UnencryptedAuth{
		smtp.PlainAuth("", mailUser, mailPWD, mailHost),
	}
	r := NewRequest(auth, mailTo, title)

	return r.SendEmail(body)
}

func main() {
	body := "have you get?<br>hahahahhah<br>ahhaha"
	_, err := Mail([]string{"xxxxx@tom.com"}, "linux test", body)

	if err != nil {
		mailErr := "mail err: " + err.Error()
		fmt.Println("fail! " + mailErr)
	}
}
