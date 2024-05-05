package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	//c := make(chan string, 3)
	//var send chan<- string = c // send-only
	//var read <-chan string = c // receive-only
	//
	//send <- "send"
	//<-read
	c := make(chan int)

	go producer(c)
	go consumer(c)

	time.Sleep(3 * time.Second)
}
