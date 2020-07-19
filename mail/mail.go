package mailUtils

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
)

func Send(host, username, password, token string, target ...string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", target...)

	m.SetHeader("Subject", "xx官方邮件")
	m.SetBody("text/html", generateBody(token))
	d := gomail.NewDialer(
		host, 465,
		username, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
