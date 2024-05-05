package main

import "fmt"

func main() {
	// 相当于char
	var c byte = 'c'
	fmt.Printf("c = %c", c)

	var c2 byte = 100
	fmt.Printf("c2 = %c", c2)

	// byte的大小 uint8 只能够展示ascii字符，其他语言的字符用rune展示
	var c1 rune = '中'
	fmt.Printf("c1 = %c", c1)
}
