package demo

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tpjg/goriakpbc"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// AliyunDemo()
	LocalDemo()
	println("-----------------------------JustDemo end>>>")
	return
}

func AliyunDemo() {
	println("//<<-------------------------AliyunDemo start-----------")
	address := "182.92.76.126:8087"
	// filePath := "/Users/Jialin/Desktop/1.png"
	// filePath := "/Users/Jialin/Downloads/好想你-i-miss-u-JoyceChu.mp4"
	filePath := "/Users/Jialin/Documents/materialUpload/webm/feelings_vp9-20130806-242.webm"
	f, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
		return

	}
	defer f.Close()
	thisb, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
		return
	}
	err = SaveFile(address, "demo", "webm1.webm", thisb)
	if err != nil {
		panic(err.Error())
		return
	}

	println("//---------------------------AliyunDemo end----------->>")
}

func LocalDemo() {
	println("//<<-------------------------LocalDemo start-----------")
	// address := "192.168.100.44:8087"
	// address := "192.168.8.3:8087"
	address := "docker.dev:8087"
	// filePath := "/Users/Jialin/Documents/sunteng/materialUpload/jpg/Qn28x6B0Dg.jpg"
	filePath := "/Users/Jialin/Documents/sunteng/materialUpload/file.mp4"
	f, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
		return

	}
	defer f.Close()
	thisb, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
		return
	}
	const bucket = "demo"
	const fn = "picd.png"
	err = SaveFile(address, bucket, fn, thisb)
	if err != nil {
		panic(err.Error())
		return
	}
	println("save to ", bucket, "/", fn)
	fmt.Printf("you may visit /buckets/{xxBucketName}/keys/{xxKeyName} from http endpoint\n")
	fmt.Printf("you may visit [/buckets/%s/keys/%s] from http endpoint\n", bucket, fn)
	// fmt.Printf("you may visit %s from http endpoint\n",/buckets/demo/keys/picd.png)

	println("//---------------------------LocalDemo end----------->>")
}

//visit data via http url:    domain:port/riak/bucketName/key
func SaveFile(address, bucketName, key string, p []byte) error {
	// var DefaultClient *riak.Client
	DefaultClient := riak.NewClient(address)
	if err := DefaultClient.Connect(); err != nil {
		panic("Cannot Connect to Riak Server:" + address)
	}
	defer DefaultClient.Close()
	bucket, err := DefaultClient.NewBucket(bucketName)
	if err != nil {
		return err
	}

	obj := bucket.NewObject(key)
	// obj.ContentType = "application/octet-stream"
	// obj.ContentType = "image/png"
	// obj.ContentType = "video/mp4"
	obj.ContentType = "video/webm"
	obj.Data = p
	return obj.Store()
}
