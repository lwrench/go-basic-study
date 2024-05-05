package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter
		if i < len(str) {
			fmt.Printf("%s%s", str[i], str[i+1])
			i += 2
		} else {
			return
		}
		number <- true
	}
}

func main() {

	go printNum()
	go printLetter()

	number <- true

	time.Sleep(2 * time.Second)
}
