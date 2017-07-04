// Package main provides ...
package demo

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type Slot struct {
	W  int `json:"W"`
	H  int `json:"H"`
	ID int `json:"ID"`
}

func (s *Slot) ToStr() string {
	return strconv.Itoa(s.W) + "x " + strconv.Itoa(s.H)
}

// type ListOfSlot []*Slot
//
// func (ls ListOfSlot) Len() int {
// 	return len(ls)
// }

type ListOfSlot []*Slot

func (ls ListOfSlot) Len() int {
	return len(ls)
}

func (ls ListOfSlot) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

func (ls ListOfSlot) Less(i, j int) bool {
	if ls[i].W > ls[j].W {
		return true
	}
	return false
}

func Unique(s []*Slot) []*Slot {
	// existMap:=make(map[string]bool)
	// ret:=make(ListOfSlot,0,len(s))
	// if existMap[v.ToStr()] ==true{
	// 	continue
	// }

	newIndex := 1
	exit := false
	for i := 1; i < len(s); i++ {
		exit = false
		for existIndex := 0; existIndex < newIndex; existIndex++ {
			if s[i].W == s[existIndex].W && s[i].H == s[existIndex].H {
				exit = true
				break
			}
		}
		if !exit {
			s[newIndex] = s[i]
			newIndex++
		}

	}
	return s[:newIndex]
}

func ReadExcelDemo() {
	println("//<<-------------------------ReadExcelDemo start-----------")
	start := time.Now()

	// xlsPath := path.Join(conf.Conf().RootDir, importIndustryConf.File)
	xlsPath := "/Users/Jialin/Downloads/media-slot.xlsx"
	xlsxFile, err := xlsx.OpenFile(xlsPath)
	if err != nil {
		panic(err.Error())
	}

	// log.Printf("len(xlsxFile.Sheets): %+v\n", len(xlsxFile.Sheets))
	// for k, v := range xlsxFile.Sheets {
	// 	fmt.Printf("%+v: %+v \n", k, v)
	// }

	sheet := xlsxFile.Sheets[0]
	RowsLen := len(sheet.Rows)
	slots := make([]*Slot, len(sheet.Rows)-1)
	// for k, v := range sheet.Rows {
	for i := 1; i < RowsLen; i++ {
		w, _ := sheet.Rows[i].Cells[6].Int()
		h, _ := sheet.Rows[i].Cells[7].Int()
		slots[i-1] = &Slot{
			ID: i - 1,
			W:  w,
			H:  h,
		}
	}
	fmt.Printf("slots:,cap(slots):%d,len(slots):%d arrd:%v \n", cap(slots), len(slots), &slots[0])
	for i := 0; i < RowsLen-1; i++ {
		// log.Printf("slots[i].W:  %+v*%d\n", slots[i].W, slots[i].H)
	}
	sort.Sort(ListOfSlot(slots))
	for k, v := range slots {
		//@toDelete
		fmt.Printf("%+v: %+v\n", k, v)
	}
	return

	retSlots := Unique(slots)
	sort.Sort(ListOfSlot(retSlots))
	retSlotsLen := len(retSlots)
	log.Printf("retSlotsLen: %+v\n", retSlotsLen)
	newSize := make([][2]int, retSlotsLen)
	for i := 0; i < retSlotsLen; i++ {
		newSize[i] = [2]int{retSlots[i].W, retSlots[i].H}
		log.Printf("new retSlots[i].W: %+v*%d\n", retSlots[i].W, retSlots[i].H)
	}
	log.Printf("newSize: %+v\n", newSize)
	bs, err := json.Marshal(newSize)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("string(bs): %+v\n", stringsToJson(string(bs)))

	// fmt.Printf("retSlots: ,cap(retSlots):%d,len(retSlots):%d arrd:%v \n", cap(retSlots), len(retSlots), &retSlots[0])

	// log.Printf("len(sheet): %+v\n", len(sheet.Rows))
	// fmt.Printf("sheet.Rows: ,cap(sheet.Rows):%d,len(sheet.Rows):%d arrd:%v \n", cap(sheet.Rows), len(sheet.Rows), &sheet.Rows[0])
	// for k, v := range sheet.Rows[0].Cells {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	//
	// println()
	// for k, v := range sheet.Rows[len(sheet.Rows)-1].Cells {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }

	// fmt.Printf("sheet.Rows: cap(sheet.Rows):%d,len(sheet.Rows):%d arrd:%v \n", cap(sheet.Rows), len(sheet.Rows), &sheet.Rows[len(sheet.Rows)-1])
	// // list := make([]*advertiser.Industry, 0, len(sheet.Rows))
	// // industryId, id := 0, 0
	// for index, row := range sheet.Rows {
	// 	log.Printf("index: %+v row:%v\n", index, row)
	// }
	//
	fmt.Printf("ReadExcelDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ReadExcelDemo end----------->>")
}
func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

func stringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}
