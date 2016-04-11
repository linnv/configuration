package main

import (
	"fmt"
	"sunteng/commons/mail"
	"time"
)

func main() {
	fmt.Println("xxxooxxx at mail demo")
	mail.InitFunc()
	time.Sleep(time.Second * 10)

}
