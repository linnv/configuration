package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

// user : example@example.com login smtp server user
// password: xxxxx login smtp server password
// host: smtp.example.com:port   smtp.163.com:25
// to: example@example.com;example1@163.com;example2@sina.com.cn;...
// subject:The subject of mail
// body: The content of mail
// mailtyoe: mail type html or text

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	// user := "jialin27@126.com"
	// host := "smtp.126.com:25"

	//not work
	user := "wujialin@sunteng.com"
	host := "smtp.exmail.qq.com:465"

	// user := "657403339@qq.com"
	// host := "smtp.qq.com:465"

	password := ""
	// to := "wujialin@sunteng.com;137140525@qq.com"
	to := "137140525@qq.com"

	subject := "Test sending email by golang"

	body := `
	<html>
	<body>
	<h3>
	"Test sending email by golang"
	</h3>
	</body>
	</html>
	`
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}
