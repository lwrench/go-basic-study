package main

import (
	"fmt"
)

func main() {
	courseSlice := make([]string, 4)
	courseSlice[0] = "go"
	courseSlice[1] = "grpc"
	courseSlice[2] = "mysql"
	courseSlice[3] = "es"
	// 会报错
	// courseSlice[4] = "gin"
	courseSlice = append(courseSlice, "gin")

	// 删除第3个元素，即mysql
	courseSliceDelete := append(courseSlice[:2], courseSlice[3:]...)
	fmt.Println(courseSliceDelete)

	newSlice := courseSlice[:3]
	// 会报错
	fmt.Println(newSlice, newSlice[3])

}
