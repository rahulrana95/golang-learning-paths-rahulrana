package main

import (
	"fmt"

	linkedList "github.com/golang-learning-paths-rahulrana/linked-list"
)

func main() {
	dummyDataIntArr := []string{"1", "2", "3", "4", "5"}
	list := &linkedList.LinkedList{}
	newList := list.CreateDummyLinkedList(dummyDataIntArr)
	newList.PrintLinkedList(true)
	newList.DeleteLinkedList()
	newList.PrintLinkedList(true)
	newList2 := list.CreateDummyLinkedList(dummyDataIntArr)
	fmt.Println("Length is", newList2.GetLengthIterative())
	fmt.Println("Length is", newList2.GetLengthRecursive())

	newList2.Delete("1")
	fmt.Println("Length is", newList2.GetLengthIterative())
	fmt.Println("Length is", newList2.GetLengthRecursive())
	newList3 := list.CreateDummyLinkedList(dummyDataIntArr)
	linkedList.MakeMiddleNodeHead(newList3)
	fmt.Println("Length is", newList3.GetLengthIterative())
	newList3.PrintLinkedList(true)

}
