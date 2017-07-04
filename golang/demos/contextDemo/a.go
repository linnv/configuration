// Package main provides ...
package demo

import (
	"context"
	"time"

	"github.com/linnv/logx"
)

// The Deadline method allows functions to determine whether they should start work at all; if too little time is left, it may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.

// Value allows a Context to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.

// the WithCancel function (described below) provides a way to cancel a new Context value.
// when calling Cancel of ctx, channel of ctx.Done() will be close, and any channel closed is readable

// After the first call, subsequent calls to a CancelFunc do nothing.
func JustDemo() {
	println("<<<JustDemo start---------------------------")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	// ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*2000)
	// 2 chan will not block return for goroutine to return
	finishChan := make(chan bool, 2)
	ctx = context.WithValue(ctx, "finishChan", finishChan)
	// ctx, cancel := context.WithCancel(ctx)
	// TimdeDelay(ctx, 1)
	// TimdeDelay(ctx, 2)
	TimdeDelay(ctx, 3)
	// go TimdeDelay(ctx, 3)
	// time.Sleep(time.Millisecond * 100)
	// cancel()
	select {
	case <-finishChan:
		println("good finish !")
		return
	case <-ctx.Done():
		println("time out", ctx.Err().Error())
		logx.Debug("ctx.Err(): %+v\n", ctx.Err())
	default:
		println("good")
		logx.Debug("good: \n")
	}

	// context.WithCancel(

	println("-----------------------------JustDemo end>>>")
	return
}

func TimdeDelay(ctx context.Context, i int) {
	switch i {
	case 2:
		finish := ctx.Value("finishChan").(chan bool)
		finish <- true
		time.Sleep(time.Second)
		return
	case 1:
		return
	default:
		time.Sleep(time.Second)
		return
	}
}
