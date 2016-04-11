package mail

import (
	"log"
	"os"
	"path"
	"time"
)

var Log *log.Logger

func Init() error {
	os.Stdout.Write(append([]byte("mail log"), '\n'))
	fileName := "/data/logs/mail_log/" + time.Now().String() + ".log"
	os.MkdirAll(path.Dir(fileName), 0777)
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("log conf error:", err.Error())
		return err
	}

	Log = log.New(logFile, "[Mail-Error]", log.Llongfile|log.LstdFlags)

	return nil
}
