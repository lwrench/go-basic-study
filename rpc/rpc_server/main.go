package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 实例化server
	listener, _ := net.Listen("tcp", ":1234")
	// 注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
