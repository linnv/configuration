package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
)

func main() {
	http.HandleFunc("/mail", mailHandler)
	port := ":8017"
	fmt.Printf("listening : %+v\n", port)
	http.ListenAndServe(port, nil)
}

type Mail struct {
	To       string `json:"To"`
	Subject  string `json:"Subject"`
	Body     string `json:"Body"`
	MailType string `json:"MailType"`
}

const (
	// user   = "biddingx@163.com"
	// passwd = "130re7q0972SDoM"
	// host   = "smtp.163.com:25"
	user   = "test27@126.com"
	passwd = "jialin27"
	host   = "smtp.126.com:25"
)

func mailHandler(w http.ResponseWriter, r *http.Request) {
	requestbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
		return
	}
	mail := Mail{}
	err = json.Unmarshal(requestbody, &mail)
	if err != nil {
		panic(err.Error())
		return
	}

	err = SendMail(mail.To, mail.Subject, mail.Body, mail.MailType)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(requestbody)
	fmt.Printf("sending mail: %+v\n", mail)
	return
}

func SendMail(to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, passwd, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	//@toDelete
	fmt.Printf("err.Error(): %+v\n", err.Error())
	return err
}
