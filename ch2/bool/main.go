package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBool, err := strconv.ParseBool("false")
	if err != nil {

	}
	fmt.Println(parseBool)
	b, err := strconv.ParseBool("1")
	if err != nil {

	}
	fmt.Println(b)

	c, err := strconv.ParseBool("0")
	if err != nil {

	}
	fmt.Println(c)

	float := strconv.FormatFloat(3.1415926, 'f', -1, 64)
	fmt.Println(float)
}
