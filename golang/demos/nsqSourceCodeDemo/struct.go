package demo

import "sync"

type handler struct {
	//excute only once, cause close a closed chan will panic
	// write to a closed chan will panic
	exitHandler sync.Once
}
