// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Property struct {
		Key [][2]string `json:"key"`
	}

	// ONE PROPERTY
	oneProp := []byte(`{"key": [["one","two"]]}`)
	var prop Property
	er := json.Unmarshal(oneProp, &prop)
	if er != nil {
		panic(er)
	} else {
		fmt.Println(prop)
	}
}
