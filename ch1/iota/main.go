package main

import "fmt"

// https://go.dev/wiki/Iota
func main() {
	const (
		Error1 = iota
		Error2
		Error3 = "error"
		Error4
		Error5 = iota
		Error6
	)

	fmt.Println(Error1, Error2, Error3, Error4, Error5, Error6)
	// iota 内部会自动计数
}
