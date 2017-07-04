// Package main provides ...
package demo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/linnv/logx"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

// type Errno uintptr

const (
	// native_client/src/trusted/service_runtime/include/sys/errno.h
	// The errors are mainly copied from Linux.
	// EPERM  Errno = 1 /* Operation not permitted */
	// ENOENT Errno = 2 /* No such file or directory */
	EPERM  = 1 /* Operation not permitted */
	ENOENT = 2 /* No such file or directory */
)

var errorstr = [...]string{
	EPERM:  "Operation not permitted",
	ENOENT: "No such file or directory",
}

var NumErrorstr = [...]string{
	1: "Operation not permitted",
	2: "No such file or directory",
}

func DisplayError() {
	for k, v := range errorstr {
		fmt.Printf("%+v: %+v\n", k, v)
	}

	fmt.Printf("  errorstr[ENOENT]: %+v\n", errorstr[ENOENT])
	fmt.Printf("  errorstr[1]: %+v\n", errorstr[1])

	for k, v := range NumErrorstr {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	//@toDelete
	fmt.Printf("NumErrorstr[1]: %+v\n", NumErrorstr[1])
}

func LinkExist(linkFN string) bool {
	fileInfo, err := os.Lstat(linkFN)
	fmt.Printf("os.ModeSymlink: %+v\n", os.ModeSymlink)

	if err != nil {
		fmt.Println(err)

		return false
		// os.Exit(1)
	}
	fmt.Printf("  fileInfo: %+v\n", fileInfo)
	println("origin file")
	return true
}

func Link() {

	// create a new symbolic or "soft" link
	originFN := "mbv.t"
	linkFN := "mbv-link.t"
	// err := os.Symlink("file.txt", "file-symlink.txt")
	err := os.Symlink(originFN, linkFN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// resolve symlinks

	fileInfo, err := os.Lstat(linkFN)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		originFile, err := os.Readlink(fileInfo.Name())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Resolved symlink to : ", originFile)
	}

}

func FilePath() {
	p := "/dd/ss/ee/x/d.g"
	var bannerName = path.Base(p)
	fmt.Printf("  bannerName: %+v\n", bannerName)
}

var (
	fh_mp4_p1 = []byte{0, 0, 0}            //[0:3]
	fh_mp4_p2 = []byte{102, 116, 121, 112} //[4:8]
	fh_mp4_p3 = []byte{0, 0}               //[12:14]
)

func BytesEqual(a, b []byte) bool {
	al := len(a)
	if al != len(b) {
		return false
	}

	for i := 0; i < al; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func FileTypeDemo() {
	println("//<<-------------------------FileTypeDemo start-----------")
	// path := "/Users/Jialin/Documents/materialUpload/BdIvTuNfDz.mp4"
	path := "/Users/Jialin/Documents/materialUpload/mp4.file"
	// path := "/Users/Jialin/Documents/materialUpload/300x250.swf"
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("打开文件失败")
		return

	}
	thisb, err := ioutil.ReadAll(f)

	if BytesEqual(thisb[0:3], fh_mp4_p1) && BytesEqual(thisb[4:8], fh_mp4_p2) && BytesEqual(thisb[12:14], fh_mp4_p3) {
		os.Stdout.Write(append([]byte("mp4 file"), '\n'))
	}

	println("//---------------------------FileTypeDemo end----------->>")
}

type bufFile struct {
	fd   *os.File
	bufw *bufio.Writer
}

func (bf *bufFile) Flush() error {
	//@TODO
	bf.bufw.Flush()
	return bf.fd.Sync()
}

func (bf *bufFile) Close() error {
	bf.Flush()
	return bf.fd.Close()
}

func (bf *bufFile) converTestWriter() {
	simpleWriter := func(w io.Writer) {
		if bw, ok := w.(*bufio.Writer); ok {
			bw.WriteString("good: io.Writer converted bufio.Writer ")
			logx.Debugln("good: io.Writer converted bufio.Writer ")
			return
		}
		switch w.(type) {
		case *bufio.Writer: //only poiter type interface work
			logx.Debugln("good: found bufio writer ")
		}
	}
	simpleWriter(bf.bufw)
}

func newBufFile() *bufFile {
	fd, err := os.OpenFile("./a.file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logx.CheckErr(err)
	bw := bufio.NewWriter(fd)
	return &bufFile{
		fd,
		bw,
	}
}

func IoerDemo() {
	println("//<<-------------------------IoerDemo start-----------")
	start := time.Now()
	bf := newBufFile()
	bf.converTestWriter()
	defer bf.Close()
	fmt.Printf("IoerDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------IoerDemo end----------->>")
}

func SaveIntSliceToFile(iSlice []int64, filePath string) error {
	// f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	l := len(iSlice)
	tmpBytes := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		tmpBytes = strconv.AppendInt(tmpBytes, iSlice[i], 10)
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	// _, err = f.WriteString(str)
	_, err = f.Write(tmpBytes)
	f.Sync()
	return nil
}

func SaveBytesToFile(bs []byte, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.Write(bs)
	f.Sync()
	return nil
}

func SaveStringToFile(str, filePath string) error {
	// f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.WriteString(str)
	f.Sync()
	return nil
}

func FileToStringDemo(filePath string) (err error, str string) {
	println("//<<-------------------------FileToStringDemo start-----------")
	// path := "/Users/Jialin/Documents/materialUpload/mp4.file"
	// path := "/Users/Jialin/Documents/materialUpload/300x250.swf"
	f, err := os.Open(filePath)
	if err != nil {
		// fmt.Errorf("打开文件失败")
		return

	}
	thisb, err := ioutil.ReadAll(f)

	println("//---------------------------FileToStringDemo end----------->>")
	return nil, string(thisb)
}

func JsonFromFileDemo(filePath string, c interface{}) (err error) {

	println("//<<-------------------------JsonFromFileDemo start-----------")
	start := time.Now()

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
	// fmt.Printf(" %v microseconds\n", time.Since(start)/1000000)
	fmt.Printf(" %v\n", time.Since(start))
	println("//---------------------------JsonFromFileDemo end----------->>")
	return
}

func FileTime() error {
	// path := "/Users/Jialin/golang/src/demo/demos/fileOperation"
	path := "/Users/Jialin/Downloads/everyone-can-use-english.epub"

	finfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	fmt.Printf("finfo: %+v\n", finfo)
	fmt.Printf("finfo.ModTime(): %+v\n", finfo.ModTime())
	fmt.Printf("size: %+v\n", b2m(finfo.Size()))

	// var prevModTime time.Time
	// if finfo.IsDir() {
	// 	os.Stdout.Write(append([]byte("dir"), '\n'))
	// } else {
	// 	modTime := finfo.ModTime()
	// 	if modTime != prevModTime {
	// 		prevModTime = modTime
	// 	}
	// }
	return nil
}

const (
	UnitKb = 1024
	UnitMb = UnitKb * UnitKb
)

func b2m(size int64) int64 {
	if size > 0 {
		return size / UnitMb
	}
	return 0
}

func b2k(size int64) int64 {
	if size > 0 {
		return size / UnitKb
	}
	return 0
}
