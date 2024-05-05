package main

import (
	"fmt"
	"go-basic-study/ch10/user" // 相较于go.mod 文件的相对路径，使用时需要使用文件中package字段的定义（course），而不是user
)

func main() {
	c := &course.Course{Name: "ssss"}
	fmt.Println(course.GetCourse(c))
}
