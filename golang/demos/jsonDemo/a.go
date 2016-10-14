package newDir

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type A struct {
	N string `json:"name"`
	B
}

type AA struct {
	N string `json:"name"`
	B *B     `json:"B struct"`
}

type AP struct {
	N string `json:"name"`
	*B
}

func (this *A) Afun() {
	println("A func pointer instance")
}

func (this A) Bfun() {
	println("A func instance")
}

type B struct {
	BN string `json:"bName"`
}

func (this *B) BFun() {
	println("b func")
}

func MarshalDemoA() (y error) {
	println("<<<MarshalDemo A---------------------------")
	a := &AP{}
	a.N = "a"
	a.B = new(B)
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}
	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo A>>>")
	return
}

func MarshalDemo() (y error) {
	println("<<<MarshalDemo---------------------------")
	a := &A{}
	a.N = "a"
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}

	aa := &AA{}
	aa.N = "aa"
	if aa.B == nil {
		aa.B = new(B)
	}
	aa.B.BN = "bn of aa"

	bs, y = json.Marshal(aa)
	if y != nil {
		return y
	}
	fmt.Printf("string(aa bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo>>>")
	return
}

func MarshalDemoAP() (y error) {
	println("<<<MarshalDemo AP---------------------------")
	a := &AP{}
	a.N = "a"
	//@TODO B is nil error occurs
	a.B = new(B)
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}
	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo AP>>>")
	return
}

func MarshalDemoP() (y error) {
	println("<<<MarshalDemo p---------------------------")
	a := A{}
	a.N = "a"
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}
	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo P>>>")
	return
}

func UnMarshalDemoAP(bs []byte) (y error) {
	a := &AP{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal AP: %+v\n", a)
	fmt.Printf("a.B: %+v\n", a.B)
	return
}

func UnMarshalDemoA(bs []byte) (y error) {
	a := AP{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal A: %+v\n", a)
	fmt.Printf("a.B: %+v\n", a.B)
	return
}
func UnMarshalDemoP(bs []byte) (y error) {
	a := &A{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal P: %+v\n", a)
	fmt.Printf("a.B: %+v\n", a.B)
	return
}

func UnMarshalDemo(bs []byte) (y error) {
	a := A{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal: %+v\n", a)
	return
}

type NN struct {
	Name string `json:"Name"`
	// Age  string `bson:"Age" json:"-"`
	// Age string `bson:"Age" json:"omitempty"`
	// Age string `bson:"Age" json:",omitempty"`
	// Age string `bson:"Age" json:",Age"`
	Age int `bson:"Age" json:",Age"`
	// Count int64  `json:",string"`
	// Count int64 `json:",string"`
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	WireUnMarshalDemoAP()
	// n := NN{Name: "xxxx", Age: "jfeifefjej", Count: 333}
	// bs, err := json.Marshal(n)
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("  string(bs): %+v\n", string(bs))
	// nn := NN{}
	// bs := `{"Name":"a","Age":"bn","Count":"r333"}`  illegal, Count must be number like string, no other characters permitted
	// bs := `{"Name":"a","Age":"bn","Count":"333"}`
	// bs := `{"Name":"a","Count":"333"}`
	// err := json.Unmarshal([]byte(bs), &nn)
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("  nn: %+v\n", nn)
	//
	// type ColorGroup struct {
	// 	ID     int
	// 	Name   string
	// 	Colors []string
	// }
	// group := ColorGroup{
	// 	ID:     1,
	// 	Name:   "Reds",
	// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	// }
	// b, err := json.Marshal(group)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)

	// a := A{}
	// a.Afun()
	// a.Bfun()
	//
	// fmt.Printf("pointer w: works\n")
	// ap := &A{}
	// ap.Afun()
	// ap.Bfun()
	// a:=1100
	// b:=555
	// var a int
	// var b int
	// a = 1000
	// b = 500
	// fmt.Printf("  a/b: %+v\n", a/b)
	// fmt.Printf("2  a/b: %+v\n", float64(a/b)
	println("-----------------------------JustDemo end>>>")
}

func LoopBasic() {
	i := 1
	i++
	fmt.Printf("  i: %+v\n", i)
	var j = 0
	// F:
	//for(i:=0;if(i<19){ do with i=0}else{break;};i++)
	for ; j < 19; j++ {
		fmt.Printf("%v: works\n", j)
		if j == 3 {
			// break F
		}
	}
	fmt.Printf("  j: %+v\n", j)
}

func BsonDemo() {
	var deleteCond = []bson.M{}
	deleteCond = append(deleteCond, bson.M{"AdSlotId": 44, "OrderId": 22})
	fmt.Printf("  deleteCond: %+v\n", deleteCond)
}

type F struct {
	W float64 `json:"W"`
	H float64 `json:"H"`
}

type FS struct {
	FS []F `json:"FS"`
}

type ImageRatio struct {
	Min WH `json:"Min "`
	Max WH `json:"Max "`
}

type WH struct {
	W float64 `json:"W"`
	H float64 `json:"H"`
}
type AT struct {
	A string `json:"A"`
	// Iter interface{} `json:"Iter"`
	Iter json.RawMessage `json:"Iter"`
}

func InterfaceDemo() {
	println("//<<-------------------------InterfaceDemo start-----------")
	start := time.Now()
	a := &AT{}
	a.A = "xxxx"
	a.Iter = &WH{1, 2}
	bs, _ := json.Marshal(a)
	fmt.Printf("string(bs): %+v\n", string(bs))
	fmt.Printf("InterfaceDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------InterfaceDemo end----------->>")
}

func LoadJsonFile(filePath string, v interface{}) error {
	fi, err := os.Stat(filePath)
	if err != nil {
		return err
	} else if fi.IsDir() {
		return errors.New(filePath + " is not a file.")
	}

	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	os.Stdout.Write(append([]byte(b), '\n'))

	err = json.Unmarshal(b, &v)
	return err
}

type T struct {
	N string `json:"N"`
	C int64  `json:"C"`

	Times time.Time
}

func TimeDemo() {
	println("//<<-------------------------TimeDemo start-----------")
	t := T{}
	t.Times = time.Now()
	tbs, err := bson.Marshal(t)
	if err != nil {
		return
	}
	fmt.Printf("  string(tbs): %+v\n", string(tbs))
	tt := T{}
	err = bson.Unmarshal(tbs, &tt)
	if err != nil {
		return
	}
	fmt.Printf("  t: %+v\n", t)
	fmt.Printf("  tt: %+v\n", tt)
	println("//---------------------------TimeDemo end----------->>")
}

func GeneratTemplate(tpl interface{}) error {
	tbs, err := bson.Marshal(tpl)
	if err != nil {
		return err
	}
	fmt.Printf("  string(tbs): %s\n", string(tbs))
	return nil
}

type Wire struct {
	A     string `json:"A"`
	Noint int    `json:"Noint"`
	In    int    `json:"In,omitempty"`
	NN    NN     `json:"NN"`
	MM    *NN    `json:"MM"`

	// In int `json:"In"`
}

func WireUnMarshalDemoAP() {
	// bs := `{"A":"xxx","In":33}`
	// a := &Wire{}
	// json.Unmarshal([]byte(bs), &a)
	// fmt.Printf("unmarshal AP: %v\n", a)
	// fmt.Printf("unmarshal AP: %+v\n", a)
	b := &Wire{}
	b.A = "xx"
	b.Noint = 3545
	ss, _ := json.Marshal(b)
	fmt.Printf("ss: %+v\n", string(ss))
	var raw json.RawMessage
	raw = ss
	c := &Wire{}
	json.Unmarshal(raw, c)
	fmt.Printf("c: %+v\n", c)

	// fmt.Printf("a.B: %+v\n", a.In)
	return
}

func JsonRaw() {
	i, j := 1, 2
	fmt.Printf("i^j: %+v %b\n", i^j, i^j)
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		//key point
		//using decoder get json struct from streamer,e.g Response.Body,Reader
		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}

		var sstatus string
		if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}
