// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/validation"
)

type User struct {
	Name string
	Age  int
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()
	u := User{"man", 40}
	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 0, 18, "age")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	// or use like this
	if v := valid.Max(u.Age, 140, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}
	// 定制错误信息
	minAge := 18
	valid.Min(u.Age, minAge, "age").Message("少儿不宜！")
	// 错误信息格式化
	valid.Min(u.Age, minAge, "age").Message("%d不禁", minAge)

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}
