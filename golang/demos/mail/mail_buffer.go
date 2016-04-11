// Package mail
package mail

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var (
	started    = false
	MailBuffer chan MailConfigChan
)

//应设置为不低于log连续出现的次数
const BufferLength = 5

// 邮件通知配置信息
type MailConfig struct {
	// 邮件来源地址
	From string
	// 邮件目的地址列表
	To []string
	// 邮件主题
	Subject string
}

type MailConfigChan struct {
	MailGroupConfig MailGroupConfig
	Content         string
}

type MailGroupConfig struct {
	SendMail bool
	Addr     string
	Groups   map[string]MailConfig
}

func InitFunc() {
	if started {
		return
	}

	started = true
	MailBuffer = make(chan MailConfigChan, BufferLength)

	Init()

	go func() {
		for {
			select {
			case mcg := <-MailBuffer:
				sendMail(mcg)
			}
		}
	}()
}

func sendMail(mcc MailConfigChan) error {
	mailConfigs := mcc.MailGroupConfig
	for topic, mail := range mailConfigs.Groups {
		msg, err := json.Marshal(map[string]interface{}{
			"To":       strings.Join(mail.To, ";"),
			"Subject":  mail.Subject,
			"Body":     topic + mcc.Content,
			"MailType": "plain",
		})
		resp, err := http.Post(mailConfigs.Addr, "application/json", bytes.NewReader(msg))
		if err != nil {
			Log.Printf("post mail router err:%s", err.Error())
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			Log.Printf("mail-sending router return err:%s", err.Error())
			return errors.New("bad return code")
		}
	}

	return nil
}
