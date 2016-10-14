// Package main provides ...
package newDir

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	ImageAnalysisDemo()
	// ffprobeDemo()
	println("-----------------------------JustDemo end>>>")
	return
}

func ImageAnalysisDemo() {
	println("//<<-------------------------ImageAnalysisDemo start-----------")
	start := time.Now()

	const dir_to_scan string = "/Users/Jialin/Desktop/img"
	files, _ := ioutil.ReadDir(dir_to_scan)
	for _, imgFile := range files {
		fmt.Printf("filepath.Join(dir_to_scan, imgFile.Name()): %+v\n", filepath.Join(dir_to_scan, imgFile.Name()))

		if reader, err := os.Open(filepath.Join(dir_to_scan, imgFile.Name())); err == nil {
			// if err != nil {
			// 	fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
			// 	return
			// }
			defer reader.Close()
			im, fm, err := image.DecodeConfig(reader)
			fmt.Printf("fm: %+v\n", fm)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
				continue
			}
			fmt.Printf("%s %d %d\n", imgFile.Name(), im.Width, im.Height)
		} else {
			fmt.Println("Impossible to open the file:", err)
		}
	}
	fmt.Printf("ImageAnalysisDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ImageAnalysisDemo end----------->>")
}

func ffprobeDemo() {
	println("//<<-------------------------ffprobeDemo start-----------")
	start := time.Now()

	const dir_to_scan string = "/Users/Jialin/Desktop/img"
	files, _ := ioutil.ReadDir(dir_to_scan)
	for _, imgFile := range files {
		fp := filepath.Join(dir_to_scan, imgFile.Name())
		fmt.Printf("filepath: %+v\n", fp)
		cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", fp)
		b, err := cmd.Output()
		if err != nil {
			fmt.Printf("err.Error(): %+v\n", err.Error())
		}
		fmt.Printf("string(b): %+v\n", string(b))
	}
	fmt.Printf("ffprobeDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ffprobeDemo end----------->>")
}
