package main

import (
	linkedList "github.com/golang-learning-paths-rahulrana/linked-list"
)

func main() {
	dummyDataIntArr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	list := &linkedList.LinkedList{}
	newList := list.CreateDummyLinkedList(dummyDataIntArr)
	newList.PrintLinkedList(true)
	newList.DeleteLinkedList()
	newList.PrintLinkedList(true)
	//newList2 := list.CreateDummyLinkedList(dummyDataIntArr)

}
