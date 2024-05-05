package main

import "fmt"

// map不是线程安全的
// sync.Map是线程安全的

func main() {
	courseMap := map[string]string{}
	courseMap["go"] = "golang"
	courseMap["rpc"] = "grpc"

	if course, ok := courseMap["go"]; !ok {
		fmt.Println("not in course map")
	} else {
		fmt.Println(course)
	}

	delete(courseMap, "rpc")
	fmt.Println(courseMap)

	// 不报错
	delete(courseMap, "dddd")
	fmt.Println(courseMap)

}
