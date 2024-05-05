package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan string
	msg = make(chan string, 1)
	msg <- "ddd"

	go func(msg chan string) {
		for value := range msg {
			fmt.Println(value)
		}
	}(msg)

	//go func(msg chan string) {
	//	for _ := range msg {
	//		msg <- "ddd"
	//	}
	//}(msg)

	msg <- "dddsss"

	time.Sleep(3 * time.Second)

}
