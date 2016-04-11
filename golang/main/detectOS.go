package main

import (
	"fmt"
	"path"
	// "path/filepath"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		println("OS is Windows")
	case "linux":
		println("OS is Linux")
	case "darwin", "freebsd":
		println("OS is Unix")
	}
	tmpP := path.Join("date", "draw.ex")
	fmt.Printf("tmpP: %+v\n", tmpP)
	// fmt.Printf("filepath.Separator: %s\n", string(filepath.Separator))
	// println(filepath.Separator)
}
