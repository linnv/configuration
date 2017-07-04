// Package main provides ...
package demo

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/linnv/logx"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	cmd := exec.Command("ls", "-all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return
	}

	// instead of using for loop, use channel to 'block' selecting
	// c.waitDone = make(chan struct{})
	// go func() {
	// 	select {
	// 	case <-c.ctx.Done():
	// 		c.Process.Kill()
	// 	case <-c.waitDone:
	// 	}
	// }()
	println("-----------------------------JustDemo end>>>")
	return
}

func ctxCommandDemo() {
	println("//<<-------------------------ctxCommandDemo start-----------")
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ls", "-all")
	// cmd := exec.CommandContext(ctx, "cat", "> /Users/Jialin/myGit/OpenDemo/golang/demos/cmdDemo/1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	normalExit := make(chan struct{})
	go func(exit chan struct{}) {
		err := cmd.Run()
		if err != nil {
			cancel()
			logx.Debug("err: %+v\n", err)
			return
		}
		close(normalExit)
	}(normalExit)
	select {
	case <-normalExit:
		logx.Debug("%s\n", "good")
		cancel()
		return
	case <-ctx.Done():
		logx.Debug("%s\n", "time out")
	}

	fmt.Printf("ctxCommandDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ctxCommandDemo end----------->>")
}
