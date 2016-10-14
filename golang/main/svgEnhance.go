// Package main provides ...
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/google/pprof/third_party/svg"
)

// svg.Message replace and add some js code to allow svg being draging,zoom in and zoom out in browser
var debug = true

func main() {
	println("//<<-------------------------svgDemo start-----------")
	start := time.Now()

	inFile := flag.String("in", "", " input file path")
	outFile := *flag.String("out", "", "output file path")
	flag.Parse()
	if *inFile == "" {
		panic("file name must be given")
	}

	name := path.Base(*inFile)
	if n := strings.LastIndex(name, ".dot"); n > 0 {
		if m := strings.LastIndex(*inFile, "/"); m > 0 {
			if outFile == "" {
				outFile = path.Join((*inFile)[:m], (name[:n] + ".svg"))
			}
		}
	}

	if debug {
		log.Printf("input File: %+v\n", *inFile)
		log.Printf("output File: %+v\n", outFile)
	}

	bs, err := ioutil.ReadFile(*inFile)
	if err != nil {
		panic(err.Error())
	}

	bsb := bytes.NewBuffer(bs)
	cmd := exec.Command("dot", "-Tsvg")
	var buf bytes.Buffer
	cmd.Stdin, cmd.Stdout, cmd.Stderr = bsb, &buf, os.Stderr
	if err := cmd.Run(); err != nil {
		if err != nil {
			panic(fmt.Errorf("Failed to execute dot. Is Graphviz installed? Error: %v", err))
		}
	}

	str := svg.Massage(buf)

	err = ioutil.WriteFile(outFile, []byte(str), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("svgDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
}
