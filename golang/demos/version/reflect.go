// Package main provides ...
package main

import (
	// "flag"
	"fmt"
	"reflect"
)

type User struct {
	Username string
}
type Admin struct {
	User
	title string
}

func main() {
	//
	// var u Admin
	// t := reflect.TypeOf(u)
	// for i, n := 0, t.NumField(); i < n; i++ {
	// 	f := t.Field(i)
	// 	fmt.Println(f.Name, f.Type)
	// }

	u := new(Admin)
	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}

}

//
// func main() {
// 	// var ip = flag.Int("flagname", 1234, "help message for flagname")
//
// 	var flagvar int
// 	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
// 	flag.Parse()
// 	// fmt.Println("ip has value ", *ip)
// 	fmt.Println("flagvar has value ", flagvar)
// }
