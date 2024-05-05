package main

import "fmt"

func main() {

	var courses [3]string
	courses[0] = "go"
	courses[1] = "gin"
	courses[2] = "grpc"

	fmt.Println(courses)

	for _, str := range courses {
		fmt.Println(str)
	}

	courses1 := [3]string{2: "grpc"}
	fmt.Println(courses1)

	courses2 := [...]string{1: "grpc"}
	fmt.Println(len(courses2))

	// 数组可以直接比较每一项
	courses3 := [...]string{"go", "gin", "grpc"}
	fmt.Println("isEqual: %v", courses == courses3)
}
