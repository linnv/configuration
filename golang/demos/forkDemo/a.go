// Package main provides ...
package newDir

import "fmt"

func uitoa(val uint) string {
	var buf [32]byte // big enough for int64
	i := len(buf) - 1
	for val >= 10 {
		buf[i] = byte(val%10 + '0')
		i--
		val /= 10
	}
	buf[i] = byte(val + '0')
	return string(buf[i:])
}

var errorstr = [...]string{
	4: "xxx4",
	8: "xxx8",
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// pid, err := syscall.ForkExec("ls", []string{"./"}, nil)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("pid: %+v\n", pid)
	// var i uint = 1
	var i int32 = 1
	// i := 1
	// fmt.Printf("i<<1: %+v\n", i<<1)
	// i <<= 1
	for i<<1 != 0 {
		i <<= 1
	}
	fmt.Printf("i: %b\n", i)
	fmt.Printf("i<<=1: %+v\n", i)

	fmt.Printf("errorstr[4]: %+v\n", errorstr[4])
	errorstrLen := len(errorstr)
	for i := 0; i < errorstrLen; i++ {
		fmt.Printf("errorstr[i]: %d->%+v\n", i, errorstr[i])
	}
	//@todoDelelte
	fmt.Printf("errorstr: %+v,cap(errorstr):%d,len(errorstr):%d  \n", errorstr, cap(errorstr), len(errorstr))

	println("-----------------------------JustDemo end>>>")
	return
}

// func syscall() {
//
// 	var rlimit, zero syscall.Rlimit
// 	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
// 	if err != nil {
// 		t.Fatalf("Getrlimit: save failed: %v", err)
// 	}
//
// 	if zero == rlimit {
// 		t.Fatalf("Getrlimit: save failed: got zero value %#v", rlimit)
// 	}
// 	fmt.Printf("rlimit: %+v\n", rlimit)
// 	fmt.Printf("zero: %+v\n", zero)
//
// 	var rusage, uzero syscall.Rusage
// 	err = syscall.Getrusage(syscall.RLIMIT_NOFILE, &rusage)
// 	if err != nil {
// 		t.Fatalf("getusage: save failed: %v", err)
// 	}
//
// 	if uzero == rusage {
// 		t.Fatalf("getusage: save failed: got zero value %#v", rlimit)
// 	}
//
// 	fmt.Printf("rusage: %+v\n", rusage)
// 	fmt.Printf("uzero: %+v\n", uzero)
// }
