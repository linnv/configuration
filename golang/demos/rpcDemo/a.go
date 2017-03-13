// Package main provides ...
package demo

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"github.com/linnv/logx"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

type ArithAddResp struct {
	Id     interface{} `json:"id"`
	Result Reply       `json:"result"`
	Error  interface{} `json:"error"`
}

func (t *Arith) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Div(args *Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = args.A / args.B
	return nil
}

func (t *Arith) Error(args *Args, reply *Reply) error {
	panic("ERROR")
}
func newRPCServer() (*rpc.Server, net.Listener) {
	newServer := rpc.NewServer()
	newServer.Register(new(Arith))

	listen, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	return newServer, listen
}

const (
	xxRpcHTTPPath = "xxrpchttp"
)

func dialRPCHTTPTimeout(network, address string, timeout time.Duration, prefixHTTPRouterForRPC string) (client *rpc.Client, err error) {
	ctx, canceler := context.WithTimeout(context.Background(), timeout)
	defer canceler()

	normal := make(chan error)
	go func() {
		client, err = rpc.DialHTTPPath(network, address, prefixHTTPRouterForRPC)
		normal <- err
	}()
	select {
	case <-normal:
		return client, err
	case <-ctx.Done():
		return nil, errors.New("dial time out")
	}
}

//@TODO bu there exist issue with timeout support, unclear what happebn behind
// func dialRPCHTTPTimeout(network, address string, timeout time.Duration, prefixHTTPRouterForRPC string) (*rpc.Client, error) {
// 	if len(strings.TrimSpace(prefixHTTPRouterForRPC)) == 0 {
// 		prefixHTTPRouterForRPC = rpc.DefaultRPCPath
// 	}
// 	var err error
// 	conn, err := net.DialTimeout(network, address, timeout)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//this tells rpc client to request the specified router for HTTP
// 	io.WriteString(conn, "CONNECT "+prefixHTTPRouterForRPC+" HTTP/1.0\n\n")
//
// 	// Require successful HTTP response
// 	// before switching to RPC protocol.
// 	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})
// 	if err == nil && resp.StatusCode == http.StatusOK {
// 		return rpc.NewClient(conn), nil
// 	}
// 	if err == nil {
// 		err = errors.New("unexpected HTTP response: " + resp.Status)
// 	}
// 	conn.Close()
// 	return nil, &net.OpError{
// 		Op:   "dial-http",
// 		Net:  network + " " + address,
// 		Addr: nil,
// 		Err:  err,
// 	}
// }

func RpcHTTPDemo() {
	newServer, listen := newRPCServer()
	//HandleHTTP of rpc server use the default HTTP mutex,we can use a new one
	//newServer.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	//any struct implement ServeHTTP is HTTP Handler, which can be used by http.Serve()
	oneHTTPMutex := http.NewServeMux()
	//path use here and used by dialRPC from client must be same
	oneHTTPMutex.Handle(xxRpcHTTPPath, newServer)
	go http.Serve(listen, oneHTTPMutex)

	// logx.EnableDevMode(false)
	logx.Debugf("lesten.Addr().String(): %+v\n", listen.Addr().String())
	// client, err := rpc.DialHTTPPath("tcp", listen.Addr().String(), xxRpcHTTPPath)
	client, err := dialRPCHTTPTimeout("tcp", listen.Addr().String(), time.Second, xxRpcHTTPPath)
	//rpc.DialHTTPPath has no timeout setting, and I just don't want to wrap this up, we can copy and update it with timeout support
	//@TODO bu there exist issue with timeout support, unclear what happebn behind
	// client, err := dialRPCHTTPTimeout("tcp", listen.Addr().String(), time.Second*2, rpc.DefaultRPCPath)
	// client, err := dialRPCHTTPTimeout("tcp", listen.Addr().String(), time.Second, xxRpcHTTPPath)
	if err != nil {
		logx.Fatalln(err)
	}

	args := &Args{1, 8}
	reply := new(Reply)
	err = client.Call("Arith.Add", args, reply)
	if err != nil {
		logx.Fatalf("Add: expected no error but got string %q", err.Error())
	}
	logx.Debugf("reply.C: %+v\n", reply.C)
	if reply.C != args.A+args.B {
		logx.Fatalf("Add: got %d expected %d", reply.C, args.A+args.B)
	}

	client.Close()
	listen.Close()
	logx.Debugf("listen: %+v\n", listen)
}

func RpcRawTCPDemo() {
	newServer, listen := newRPCServer()
	go newServer.Accept(listen)

	args := &Args{1, 8}
	reply := new(Reply)
	client, err := rpc.Dial("tcp", listen.Addr().String())
	if err != nil {
		panic(err)
	}

	err = client.Call("Arith.Add", args, reply)
	if err != nil {
		logx.Fatalf("Add: expected no error but got string %q", err.Error())
	}
	logx.Debugf("reply.C: %+v\n", reply.C)
	if reply.C != args.A+args.B {
		logx.Fatalf("Add: got %d expected %d", reply.C, args.A+args.B)
	}

	client.Close()
	listen.Close()

	return
}

func Init() {
	rpc.Register(new(Arith))
}

func RpcDemo() {
	println("//<<-------------------------RpcDemo start-----------")
	Init()
	start := time.Now()
	//test code from stander lib
	cli, srv := net.Pipe()
	go jsonrpc.ServeConn(srv)

	client := jsonrpc.NewClient(cli)
	defer client.Close()

	// Synchronous calls
	args := &Args{1, 8}
	reply := new(Reply)
	err := client.Call("Arith.Add", args, reply)
	if err != nil {
		logx.Fatalf("Add: expected no error but got string %q", err.Error())
	}
	logx.Debugf("reply.C: %+v\n", reply.C)
	if reply.C != args.A+args.B {
		logx.Fatalf("Add: got %d expected %d", reply.C, args.A+args.B)
	}
	fmt.Printf("RpcDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------RpcDemo end----------->>")
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}
