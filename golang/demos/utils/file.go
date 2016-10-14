// Package utility
package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func SaveAnyToFile(str, filePath string) error {
	//@TODO case: bytes, int,string,float
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.WriteString(str)
	f.Sync()
	return nil
}

func SaveStringToFile(str, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.WriteString(str)
	f.Sync()
	return nil
}

func SaveBytesToFile(bs []byte, filePath string) error {
	// f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.Write(bs)
	f.Sync()
	return nil
}

func FileToStringDemo(filePath string) (err error, str string) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	thisb, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	return nil, string(thisb)
}

func JsonFromFileDemo(filePath string, c interface{}) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	var bs []byte

	bs, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &c)
	return
}
