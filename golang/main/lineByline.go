// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var logFile map[int]string

func ReadLineByLine(filepath string) {
	inputFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	// Closes the file when we leave the scope of the current function,
	// this makes sure we never forget to close the file if the
	// function can exit in multiple places.
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// scanner.Scan() advances to the next token returning false if an error was encountered
	i := 0
	logFile = make(map[int]string, 0) //len
	for scanner.Scan() {
		i++
		logFile[i] = scanner.Text()
		// fmt.Println(i, ":", scanner.Text())

	}

	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}
func main() {
	ReadLineByLine("./mongo.go")

	for k, v := range logFile {
		fmt.Printf("%+v: %+v\n", k, v)
	}
}
