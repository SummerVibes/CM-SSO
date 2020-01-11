package utils

import (
	"github.com/spf13/viper"
	"log"
	"net/smtp"
	"strings"
)

func SendMail(to []string,body string) error {
	username := viper.GetString("mail.auth.username")
	auth := smtp.PlainAuth("",
		username,
		viper.GetString("mail.auth.password"),
		viper.GetString("mail.auth.host"))
	from := viper.GetString("mail.from")
	subject := "请点击下方链接完成注册"
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + from +
		"<" + username + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(viper.GetString("mail.addr"), auth, username, to, msg)
	if err != nil {
		log.Printf("send mail error: %v", err)
	}
	return err
}

