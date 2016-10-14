package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"
)

// import "unsafe"

type uintreg uint32
type intptr int32 // TODO(rsc): remove

const ptrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
// const regSize = 4 << (^uintreg(0) >> 63) // unsafe.Sizeof(uintreg(0)) but an ideal const

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Printf("os.Getpid(): %+v\n", os.Getpid())
	fmt.Printf("syscall.Getppid(): %+v\n", syscall.Getppid())
	fmt.Printf("syscall.Getpid(): %+v\n", syscall.Getpid())
	// Declarations for runtime services implemented in C or assembly.

	// Max stack size is 1 GB on 64-bit, 250 MB on 32-bit.
	// Using decimal instead of binary GB and MB because
	// they look nicer in the stack overflow failure message.
	if ptrSize == 8 {
		println("64 bit system")
		// maxstacksize = 1000000000
	} else {
		println("32 bit system")
		// maxstacksize = 250000000
	}
	time.Sleep(time.Second * 90)

}
