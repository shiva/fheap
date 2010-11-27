package main

import (
	"container/list"
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Testing Lists")
	testLists()
	
	
}

// testing lists

func testLists() {
	
	// create list
	fmt.Println("Creating list ...")
	l := list.New()
	for i := 1; i <=10; i++ {
		l.PushBack(i)
	}
	
	fmt.Println("Printing list:")
	// iterate and print
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

