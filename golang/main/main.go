package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var maxSize int = 100000

func main() {

	f, _ := os.OpenFile("./tmp2.json", os.O_RDONLY, 0660)
	reader := bufio.NewReader(f)

	i := 0
	for {
		str, err := reader.ReadString('\n')
		i++
		if err == io.EOF {
			break
		}
		if i == 1 {
			continue
		}
		tmpb := []byte(str)
		for i, v := range tmpb {
			fmt.Printf("%d->%d ", i, v)
		}
		println()
		// fmt.Printf("%dbefore str: %+v\n", i, str)
		strings.TrimSuffix("xxxx,xxxoo,", "<,")
		replacer := strings.NewReplacer("\n", "")
		// fmt.Println("before  ->", str)
		str = str[:len(str)-3]
		tmpx := []byte(str)
		for i, v := range tmpx {
			fmt.Printf("%d=>%d ", i, v)
		}
		println()
		println()
		// str = strings.Trim(replacer.Replace(str), ",")
		// fmt.Println("after  ->", str)
		// fmt.Printf("%dafter str: %+v\n", i, str)
		// fmt.Println("after-all  ->", str+"fuck")

		var mp = map[string]string{}
		err = json.Unmarshal([]byte("{"+str+"}"), &mp)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
