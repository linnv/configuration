// Package main provides ...
package newDir

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// as := []int{1, 2}
	// bs := make([]int, 3)
	// // copy as to bs
	// //target source
	// //to   from
	// copy(bs, as)
	// fmt.Printf("bs: %+v\n", bs)
	// as = bs
	// // bs = as[0:len(as)]
	// fmt.Printf("after bs: %+v\n", bs)
	// fmt.Printf("bs: %+v,cap(bs):%d,len(bs):%d  \n", bs, cap(bs), len(bs))
	// //@todoDelelte
	// fmt.Printf("as: %+v,cap(as):%d,len(as):%d  \n", as, cap(as), len(as))
	// fmt.Printf("bs[1:]: %+v\n", bs[1:])

	var a int32 = 44
	br := atomic.CompareAndSwapInt32(&a, a, 1)
	fmt.Printf("br: %+v\n", br)
	fmt.Printf("a: %+v\n", a)

	// path := "/Users/Jialin/golang/src/demo"
	// f, err := os.Open(path)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// // 遍历当前文件夹
	// fs, err := f.Readdirnames(0)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// fmt.Printf("  fs: %+v\n", fs)
	// s := filepath.Dir(fp)
	// fmt.Printf("  s: %+v\n", s)
	println("-----------------------------JustDemo end>>>")
	return
}

func WriteStrDemo() {
	fmt.Printf("  byte(1): %+v\n", byte(1))
	fmt.Printf("  byte(0): %+v\n", byte(0))
	s := " world"
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，以便于打印
	buf.WriteString(s)        //将s这个string写到buf的尾部
	fmt.Println(buf.String()) //打印 hello world
}

func WriteByteDemo() {
	println("<<<WriteByteDemo start---------------------------")
	var s byte = '!'
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) //buf.String()方法是吧buf里的内容转成string，以便于打印
	buf.WriteByte(s)          //将s这个string写到buf的尾部
	fmt.Println(buf.String()) //打印 hello!
	println("-----------------------------WriteByteDemo end>>>")

}

func WriteBufferDemo() {
	println("//<<-------------------------WriteBufferDemo start-----------")
	// imgMaterials := []string{"jj", "ee", "pppkje"}
	// bannerNames := bytes.NewBufferString("[\"")
	// for k, v := range imgMaterials {
	// 	if k != len(imgMaterials)-1 {
	// 		bannerNames.WriteString(v + "\",\"")
	// 	} else {
	// 		bannerNames.WriteString(v + "\"")
	// 	}
	// }
	// bannerNames.WriteString("]")
	// fmt.Printf("  s: %s\n", bannerNames.String())

	bufCap := 1 << 20
	buffer := bytes.NewBuffer(make([]byte, 0, bufCap))
	fmt.Printf("buffer.Available(): %+v\n", buffer.Len())
	buffer.WriteString("a")
	buffer.WriteString("b")
	fmt.Printf("buffer.Available(): %+v\n", buffer.Len())
	c, err := buffer.ReadByte()
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Printf("c: %+v\n", c)
	fmt.Printf("buffer.Available(): %+v\n", buffer.Len())
	println("//---------------------------WriteBufferDemo end----------->>")
}

// 文件日志操作类
type FileLog struct {
	mu       sync.Mutex
	file     *os.File
	bufFile  *bufio.Writer
	lastTime time.Time
	name     string
	buf      []byte
}

const (
	BUF_SIZE = 32 << 10 //32kb
)

func FileBufferDemo() {
	println("//<<-------------------------FileBufferDemo start-----------")
	flog := &FileLog{}
	logPath := "/Users/Jialin/golang/src/demo/bytesDemo/log.t"
	var err error
	flog.file, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err.Error())
		return
	}
	flog.bufFile = bufio.NewWriterSize(flog.file, BUF_SIZE)
	flog.bufFile.WriteString("jialin")
	v := flog.bufFile.Available()
	fmt.Printf(" 1 v: %+v\n", v)
	flog.bufFile.WriteString("jialin")
	v = flog.bufFile.Available()
	fmt.Printf(" 2 v: %+v\n", v)
	flog.bufFile.Flush()

	flog.buf = make([]byte, 0, BUF_SIZE)
	v = flog.bufFile.Available()
	fmt.Printf("  v: %+v\n", v)
	flog.bufFile.Flush()
	flog.file.Close()
	// flog.buf.WriteString("jialin")
	println("//---------------------------FileBufferDemo end----------->>")
}

func ReadBytesDemo() {
	println("//<<-------------------------ReadBytesDemo start-----------")
	start := time.Now()
	f, _ := os.Open("./log.t")
	// b := make([]byte, 3) //three bytes per chinese character
	b := make([]byte, 11)
	n, _ := f.ReadAt(b, 0)
	fmt.Println(n)
	fmt.Println(string(b[:n]))
	for i := 0; i < len(b); i++ {
		println(b[i])
	}
	fmt.Printf(" %v microseconds\n", time.Since(start)/1000)
	println("//---------------------------ReadBytesDemo end----------->>")
}
