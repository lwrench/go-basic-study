package main

import "fmt"

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("error type")
	}

}
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add("1", "2"))
}
