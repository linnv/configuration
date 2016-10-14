// Package main provides ...
package newDir

import (
	"fmt"
	"os"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	dns := StrFromEnv()
	fmt.Printf("dns: %+v\n", dns)
	println("-----------------------------JustDemo end>>>")
	return
}

var (
	user      string
	pass      string
	prot      string
	addr      string
	dbname    string
	dsn       string
	netAddr   string
	available bool
)

func StrFromEnv() string {
	// get environment variables
	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}
	user = env("MYSQL_TEST_USER", "root")
	pass = env("MYSQL_TEST_PASS", "defalt password")
	prot = env("MYSQL_TEST_PROT", "tcp")
	addr = env("MYSQL_TEST_ADDR", "localhost:3306")
	dbname = env("MYSQL_TEST_DBNAME", "gotest")
	netAddr = fmt.Sprintf("%s(%s)", prot, addr)
	time.Now().Unix()
	return fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
	// c, err := net.Dial(prot, addr)
	// if err == nil {
	// 	available = true
	// 	c.Close()
	// }
}
