package main

import "fmt"

func main() {
	done := false

	i := 1
	go func() {
		done = true
		if i > 29 {
			println("exit")
			return
		}
	}()

	// go func() {
	for !done {
		// It'll be a problem as long as it contains code that doesn't trigger the scheduler execution
		//上面的goroutine并不会马上执行，因这里堵塞了进程，当发生system call,lock operation, channel operation,gc,执行goroutine才会解除堵塞,即trigger the scheduler execution
		//这里还是会一直堵塞进程
		i++

		// Gosched yields the processor, allowing other goroutines to run.  It does not
		// suspend the current goroutine, so execution resumes automatically.
		// runtime.Gosched()

		//增加这一句并不会堵塞,有系统调用
		// println("for loop")

		//增加这一句并不会堵塞,有系统调用
		// time.Sleep(time.Second)

	}
	// }()

	fmt.Println("done!")
}
