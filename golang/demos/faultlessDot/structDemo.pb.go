// Package main provides ...
package faultlessDot

import "encoding/json"

type A struct {
	N string `json:"name"`
}

func MarshalDemo(a A) string {
	println("<<<MarshDemo---------------------------")
	ret, y := json.Marshal(a)
	if y != nil {
		println("error", y.Error())
	}
	println("-----------------------------MarshDemo>>>")
	return string(ret)
}
