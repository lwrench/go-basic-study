package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "china 中国"
	bytes := []rune(name)
	fmt.Println(len(bytes), len(name))

	courseName := "go \"课程\""
	fmt.Println(courseName)

	// string builder性能高， Sprintf可读性好
	builder := strings.Builder{}
	builder.WriteString("用户名：")
	builder.WriteString("lwrench")
	re := builder.String()
	fmt.Println(re)

	fmt.Println(strings.Trim("#&hello d#", "&#"))
}
