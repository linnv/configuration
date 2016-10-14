// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"demos/utils"
)

type FW struct {
	file *os.File
}

func (fw *FW) Open() (err error) {
	fw.file, err = os.OpenFile("/Users/Jialin/golang/src/demos/demos/stdIODemo/t.t", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	//fd from open()  is read only
	// fw.file, err = os.Open("/Users/Jialin/golang/src/demos/demos/stdIODemo/t.t")
	return
}

func (fw *FW) Write(bs []byte) (int, error) {
	fmt.Printf("writing:  to file%+v\n", fw.file.Name())
	return fw.file.Write(bs)
}

func (fw *FW) Close() error {
	defer fw.file.Sync()
	return fw.file.Close()
}
func pipDemo() {
	println("//<<-------------------------pipDemo start-----------")
	start := time.Now()
	// r, w := io.Pipe()
	fmt.Printf("pipDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------pipDemo end----------->>")
}

func uuidDemo() {
	println("//<<-------------------------uuidDemo start-----------")
	start := time.Now()
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("string(out): %+v\n", string(out))
	fmt.Printf("uuidDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------uuidDemo end----------->>")
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	uuidDemo()
	// cmd := exec.Command("/Users/Jialin/golang/src/demos/demos/stdIODemo/run")
	// // cmd.Stdout = os.Stdout
	//
	// file, err := os.OpenFile("/Users/Jialin/golang/src/demos/demos/stdIODemo/t.t", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	// filein, err := os.OpenFile("/Users/Jialin/golang/src/demos/demos/stdIODemo/in.t", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	// defer file.Close()
	// defer filein.Close()
	// utils.CheckErr(err)
	// cmd.Stdin = filein
	// cmd.Stdout = os.Stdout
	// // cmd.Stdout = file
	//
	// // var b bytes.Buffer
	// // cmd.Stdout = &b
	// err = cmd.Run()
	// utils.CheckErr(err)

	// file := &FW{}
	// err = file.Open()
	// defer file.Close()
	// utils.CheckErr(err)
	// _, err = fmt.Fprintf(file, "%s", b.String())
	// utils.CheckErr(err)

	// WriterReader()
	println("-----------------------------JustDemo end>>>")
	return
}

func WriterReader() {
	println("//<<-------------------------Demo start-----------")

	// reader := strings.NewReader("jialinwu")
	// bs := make([]byte, 8)
	// // bs := make([]byte, 7)
	// n, err := reader.ReadAt(bs, 0)
	// if err != nil {
	// 	if err.Error() == "EOF" {
	// 		os.Stdout.Write(append([]byte("eof"), '\n'))
	// 	}
	// 	fmt.Printf("err.Error(): %+v\n", err.Error())
	// }
	// // utils.CheckErr(err)
	// fmt.Printf("bs: %+v\n", string(bs))
	// fmt.Printf("n: %+v\n", n)
	file, err := os.Open("/Users/Jialin/golang/src/demos/demos/stdIODemo/t.t")
	utils.CheckErr(err)
	stat, err := file.Stat()
	utils.CheckErr(err)
	total := stat.Size()
	fmt.Printf("file size: %+v\n", total)
	for i := 1; i < 50; i++ {
		start := time.Now()
		var readSize int64 = 0
		len := i * 30
		output, buffer := make([]byte, 0, total), make([]byte, len)
		for readSize < total {
			n, err := file.ReadAt(buffer, readSize)
			if err != nil && err.Error() == "EOF" {
				// println("reach end of file")
			}
			readSize += int64(n)
			output = append(output, buffer...)
		}
		// fmt.Printf("%s\b", string(output))
		fmt.Printf("buffer length %d,Demo costs  %d millisecons actually %v\n", len, time.Since(start).Nanoseconds()/1000000, time.Since(start))
		readSize = 0
	}
	println("//---------------------------Demo end----------->>")
	return
}
