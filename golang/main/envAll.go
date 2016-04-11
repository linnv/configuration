// Package main provides ...
package main

import (
	"fmt"
	"os"
)

var (
	user      string
	pass      string
	prot      string
	addr      string
	dbname    string
	dsn       string
	netAddr   string
	gopath    string
	available bool
)

func EnvAll() {
	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}

	user = env("MYSQL_TEST_USER", "root")
	gopath = env("GOPATH", "root")
	pass = env("MYSQL_TEST_PASS", "")
	prot = env("MYSQL_TEST_PROT", "tcp")
	addr = env("MYSQL_TEST_ADDR", "localhost:3306")
	dbname = env("MYSQL_TEST_DBNAME", "gotest")
	netAddr = fmt.Sprintf("%s(%s)", prot, addr)
	dsn = fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
	fmt.Printf("dsn: %+v\n", dsn)
	fmt.Printf("netAddr: %+v\n", netAddr)
	fmt.Printf("gopath: %+v\n", gopath)
}

func main() {
	fmt.Println("env")
	EnvAll()
}
