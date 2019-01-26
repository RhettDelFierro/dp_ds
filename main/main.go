package main

import (
	"github.com/RhettDelFierro/designPatterns/datastructures/linkedlist/mine"
	"fmt"
)
func main() {
	linkedList := &ds.LinkedList{}
	linkedList.InsertFirst("test")
	fmt.Printf("%+v\n", *linkedList.Head)
	linkedList.InsertFirst("twotest")
	fmt.Printf("%+v\n", *linkedList.Head.Next)
}