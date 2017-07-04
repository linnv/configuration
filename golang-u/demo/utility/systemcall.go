package utility

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 等待信号
// 如果信号参数为空，则会等待常见的终止信号
// SIGHUP 1 A 终端挂起或者控制进程终止
// SIGINT 2 A 键盘中断（如break键被按下）
// SIGQUIT 3 C 键盘的退出键被按下
// SIGKILL 9 AEF Kill信号
// SIGTERM 15 A 终止信号
// SIGCHLD 20,17,18 B 子进程结束信号
// SIGSTOP 17,19,23 DEF 终止进程
func SystemSignalWaiter(sig ...os.Signal) os.Signal {
	fmt.Println("do something to exit")
	c := make(chan os.Signal, 1)
	if len(sig) == 0 {
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGCHLD, syscall.SIGSTOP)
	} else {
		signal.Notify(c, sig...)
	}
	return <-c
}
