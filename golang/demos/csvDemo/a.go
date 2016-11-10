// Package main provides ...
package newDir

import (
	"bytes"
	"encoding/csv"
	"log"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Ken33", "2", "44"},
		{"Robert", "Griesemer", "gri"},
	}
	// file, err := os.Create("./result.csv")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer file.Close()
	// w := csv.NewWriter(os.Stdout)
	// var bf bytes.Buffer

	// b := &bytes.Buffer{}
	b := bytes.Buffer{}
	w := csv.NewWriter(&b)
	// bf := bufio.NewWriter(nil)
	// w := csv.NewWriter(bf)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	log.Printf("b: %+v\n", b.String())

	println("-----------------------------JustDemo end>>>")
	return
}
