// Package main provides ...
package newDir

import (
	"bytes"
	"log"
	"net/smtp"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

// 邮件通知配置信息
type MailConfig struct {
	// 邮件来源地址
	From string
	// 邮件目的地址列表
	To []string
	// 邮件转发服务器
	SmtpServer string
	// 用户认证服务器
	AuthServer string
	// 用户名
	UserName string
	// 密码
	Password string
	// 邮件主题
	Subject string
	// 邮件附加头
	Extra string
}

type MailGroupConfig struct {
	Addr   string
	Groups map[string]MailConfig
}

func SendMail() (err error) {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mail.sunteng.com:587")
	if err != nil {
		log.Print("dial faild")
		return err
	}
	if err = c.StartTLS(nil); err != nil {
		log.Print("tls faild")
		return
	}
	auth := smtp.PlainAuth(
		"",
		"dsp_masky@sunteng.com",
		"123456",
		"mail.sunteng.com",
	)
	if err = c.Auth(auth); err != nil {
		log.Print("auth faild")
		return
	}
	// Set the sender and recipient.
	if err = c.Mail("dsp_masky@sunteng.com"); err != nil {
		log.Print("mail faild")
		return
	}
	if err = c.Rcpt("liuwei@sunteng.com"); err != nil {
		log.Print("rcpt faild")
		return
	}
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Print("data faild")
		return err
	}
	defer wc.Close()
	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		return
	}
	return
}
