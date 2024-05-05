package main

import "fmt"

func main() {
	const (
		a = 1
		b
		c = "asd"
		d
	)

	fmt.Println(a, b, c, d)
	// a=1,b=1,c="asd",d="asd"
}
