// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"os"
	"path"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2009/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func LogToFile() {
	fileName := "./log/demo.log"

	fmt.Printf("  fileName: %+v\n", fileName)
	os.MkdirAll(path.Dir(fileName), 0777)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// os.Stderr.Write(append([]byte(""), "This is a test log entry"...))

	log.SetOutput(f)
	// log.SetFlags(log.Lshortfile | LstdFlags)
	log.SetFlags(log.Llongfile | LstdFlags)
	log.SetPrefix("[Error]")

	log.Println("This is a test log entry")

	// fmt.Printf("  Ldate: %+v\n", Ldate)
	// fmt.Printf("  Ltime: %+v\n", Ltime)
	// fmt.Printf("  Lmicroseconds: %+v\n", Lmicroseconds)
	// fmt.Printf("  Lshortfile: %+v\n", Lshortfile)
	// fmt.Printf("  LstdFlags: %+v\n", LstdFlags)
	//
	// fmt.Printf("Ltime: %v\n", strconv.FormatInt(Ltime, 2))
	// fmt.Printf("Ldate: %v\n", strconv.FormatInt(Ldate, 2))
	// fmt.Printf("LstdFlags: %v\n", strconv.FormatInt(LstdFlags, 2))
	// fmt.Printf("Lmicroseconds: %v\n", strconv.FormatInt(Lmicroseconds, 2))
	// fmt.Printf("Lshortfile: %v\n", strconv.FormatInt(Lshortfile, 2))
	// fmt.Printf("Llongfile: %v\n", strconv.FormatInt(Llongfile, 2))
	// // fmt.Printf("  b: %b\n", Lmicroseconds)
	// os.Stdout.Write(append([]byte("hehfehfe"), '\n'))
	// os.Stdout.WriteString("dd good")
}
