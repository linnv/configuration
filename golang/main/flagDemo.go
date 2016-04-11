package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	excelFileName := flag.String("file", "/Users/Jialin/Downloads/siteCategory1.0.xlsx", "excel file to be imported")

	backupOption := flag.Bool("dump", false, "backup catetory collection before importation")
	ownerId := flag.Int("ownerId", 0, "ownerId")
	flag.Parse()

	fileName := "./demo.log"
	fmt.Printf("  fileName: %+v\n", fileName)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
	log.Printf("  excelFileName: %+v\n", *excelFileName)
	log.Printf("  backupOption: %+v\n", *backupOption)
	log.Printf("  ownerId: %+v\n", *ownerId)
}

// func main() {
// 	excelFileName := flag.String("file", "/Users/Jialin/Downloads/siteCategory1.0.xlsx", "excel file to be imported")
//
// 	backupOption := flag.Bool("dump", false, "backup catetory collection before importation")
// 	ownerId := flag.Int("ownerId", 0, "ownerId")
// 	flag.Parse()
//
// 	fileName := "./demo.log"
// 	fmt.Printf("  fileName: %+v\n", fileName)
// 	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		fmt.Printf("error opening file: %v", err)
// 		return
// 	}
// 	defer f.Close()
//
// 	log.SetOutput(f)
// 	log.Println("This is a test log entry")
// 	log.Printf("  excelFileName: %+v\n", *excelFileName)
// 	log.Printf("  backupOption: %+v\n", *backupOption)
// 	log.Printf("  ownerId: %+v\n", *ownerId)
// }
