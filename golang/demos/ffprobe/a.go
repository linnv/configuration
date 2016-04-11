// Package main provides ...
package newDir

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")

	// filesPath := []string{"//Users/Jialin/Downloads/mmm.mp4", "//Users/Jialin/Downloads/small.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-171.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-172.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-242.webm", "/Users/Jialin/Downloads/out9.webm", "/Users/Jialin/Downloads/out8.webm"}
	// filesPath := []string{"//Users/Jialin/Downloads/mmm.mp4", "//Users/Jialin/Downloads/small.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-171.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-172.mp4", "/Users/Jialin/Downloads/feelings_vp9-20130806-242.webm", "/Users/Jialin/Downloads/out9.webm", "/Users/Jialin/Downloads/out8.webm"}
	// filesPath := []string{"//Users/Jialin/Downloads/mmm.webm", "//Users/Jialin/Downloads/small.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-171.webm", "/Users/Jialin/Downloads/feelings_vp9-20130806-172.mp4", "/Users/Jialin/Downloads/feelings_vp9-20130806-242.webm", "/Users/Jialin/Downloads/out9.webm", "/Users/Jialin/Downloads/out8.webm"}
	filesPath := []string{"/Users/Jialin/Downloads/压缩后的MP4文件.webm"}
	// BL_BUF_SIZE := 2 << 11 //2kb
	// bs := bytes.NewBuffer(make([]byte, 0, BL_BUF_SIZE))
	// bsHeader := bytes.NewBuffer(make([]byte, 0, BL_BUF_SIZE))
	for k, v := range filesPath {
		// bs.WriteString(v)
		// bs.WriteByte('\n')
		//@toDelete
		fmt.Printf(" %dth v: %+v\n", k, v)
		tbs := GetFileHeaderDemo(v)
		//@toDelete
		fmt.Printf("  tbs: %+v\n", tbs)
		// bs.Write(tbs)
		// bs.WriteByte('\n')
		// bs.WriteString(ffprobeDemo(v))
		// bs.WriteByte('\n')

		// bsHeader.WriteString(v)
		// bsHeader.WriteByte('\n')
		// bsHeader.Write(tbs)
		// bsHeader.WriteByte('\n')
		// bsHeader.WriteByte('\n')

	}
	//@toDelete
	// fmt.Printf("  bsHeader.Bytes(): %+v\n", bsHeader.Bytes())
	// filePath := "/Users/Jialin/video.mp4"
	// fileSavePathAll := "/Users/Jialin/webmAll.log"
	// fileSavePathHeader := "/Users/Jialin/webmHeader.log"
	// utility.SaveBytesToFile(bs.Bytes(), fileSavePathAll)
	// utility.SaveBytesToFile(bsHeader.Bytes(), fileSavePathHeader)
	println("-----------------------------JustDemo end>>>")
	return
}

func GetVideoJPEGShot(in, out string, skipSecond int) error {
	//ffmpeg -ss 1 -r 1 -i test.mp4 -vframes 1 -f mjpeg -y test.jpg
	cmd := exec.Command("ffmpeg", "-ss", strconv.Itoa(skipSecond), "-r", "1", "-i", in, "-vframes", "1", "-f", "mjpeg", "-y", out)
	_, err := cmd.Output()
	//@toDelete
	fmt.Printf("GetVideoJPEGShot   err: %+v\n", err)
	return err
}

func GetFileHeaderDemo(filePath string) []byte {
	bs := make([]byte, 20)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	if _, err = f.Read(bs); err != nil {
		panic(err.Error())
	}
	f.Close()
	return bs
}

func ffprobeDemo(filePath string) string {
	println("//<<-------------------------ffprobeDemo start-----------")
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filePath)
	b, err := cmd.Output()
	// os.Remove(file)
	if err != nil {
		panic(err.Error())
	}
	println("//---------------------------ffprobeDemo end----------->>")
	return string(b)
}
