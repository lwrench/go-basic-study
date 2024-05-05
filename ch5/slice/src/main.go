package main

import (
	"fmt"
	"strconv"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}
func main() {
	courses := []string{
		"go", "grpc", "gin",
	}
	courses1 := courses[1:2]
	max, min := 1, 2
	fmt.Printf("courses1:%v\r\n", courses1)
	fmt.Printf("courses1:%v\r\n", courses1[max-min])

	printSlice(courses)
	fmt.Println(courses)
}
