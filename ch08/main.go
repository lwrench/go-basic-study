package main

import "fmt"

//func swap(num1, num2 *int) {
//	num1, num2 = num2, num1
//}

func swap(num1, num2 *int) {
	// 用临时变量temp存储num1地址的值
	temp := *num1
	// 将num1的值改为num2的值
	*num1 = *num2
	// 将num2的值改为临时变量存储的值
	*num2 = temp
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	// 交换失败
	fmt.Println(a, b)
}
