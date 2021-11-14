package main

import (
	"fmt"

	leetcodeContest267 "github.com/golang-learning-paths-rahulrana/leetcode-contests"
	linkedList "github.com/golang-learning-paths-rahulrana/linked-list"
)

func main() {
	list := linkedList.LinkedList{}
	arr := []int{1, 2, 3, 4}
	list.Add("1")
	list.Add("2")
	list.Add("3")
	list.Add("4")
	// list.delete("1")
	// list.delete("2")
	// list.delete("3")
	// list.delete("4")
	// list.reverseLinkedList()
	list.PrintReverseOrder()
	fmt.Println("---")
	list.PrintLinkedList()
	fmt.Println("Welcome")

	leetcodeContest267.TimeRequiredToBuy(arr, 2)
}
