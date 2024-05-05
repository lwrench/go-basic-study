package main

import (
	"container/list"
	"fmt"
)

// 双向链表
func main() {
	//mylist := list.List{}
	mylist := list.New()

	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("mysql")

	mylist.PushFront("java")

	fmt.Println(mylist)

	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for i := mylist.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	i := mylist.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "grpc" {
			break
		}
	}

	mylist.InsertBefore("pgsql", i)
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	mylist.Remove(i)

}
