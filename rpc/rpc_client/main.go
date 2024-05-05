package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic("error in connect")
	}
	var reply string
	// 无法做到像本地调用的效果，必须知道服务端提供的函数名称，然后使用Call调用
	// 期望的是Hello()函数能够像写才client端的本地函数一样，直接通过client.Hello("lwrench", &reply)的方式调用
	err = client.Call("HelloService.Hello", "lwrench", &reply)
	if err != nil {
		panic("error in connect")
	}
	fmt.Println(reply)
}
