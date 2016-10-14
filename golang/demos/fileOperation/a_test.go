// Package main provides ...
package newDir

import "testing"

func TestJustDemo(t *testing.T) {
	// JustDemo()
	// DisplayError()
	ReaddirsDemo()
	// Link()
	// FileTypeDemo()
	// basic := "/Users/Jialin/web/demo/chart/samples/"
	// absPath := basic + "line.html"
	// _, str := FileToStringDemo(absPath)
	//
	// replacer := strings.NewReplacer(
	// 	"{{dateToFill}}", "xnefnen0000x",
	// )
	// str = replacer.Replace(str)
	// str := "333"
	// SaveStringToFile(str, basic+"demoxx.html")
	// SaveStringToFile(str, basic+"tmp.html")

	// originFN := "mbv.t"
	// LinkExist(originFN
	// FilePath()
	// linkFN := "mbv-link.t"
	// LinkExist(linkFN)

	// is := []int64{1, 2, 4, 34, 3}
	// filepath := "/Users/Jialin/Desktop/ints.log"
	// err := SaveIntSliceToFile(is, filepath)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// FileTime()

	// filepath := "/Users/Jialin/Desktop/ints.log"
	// c := &CC{}
	//
	// err := JsonFromFileDemo(filepath, &c)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// fmt.Printf("  c: %+v\n", len(c.Index))
}

type CC struct {
	Index []int `json:"num"`
}
