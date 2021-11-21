package linkedList

import "fmt"

type Node struct {
	value string
	next  *Node
	// will be used for doubly linked list
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) Add(value string) {
	newNode := &Node{value: value}
	tempNode := list.head

	if list.head == nil {
		list.head = newNode
		list.tail = list.head
	} else {
		for tempNode.next != nil {
			tempNode = tempNode.next
		}
		tempNode.next = newNode
		list.tail = newNode
	}
}

func (list *LinkedList) PrintLinkedList(printExtra bool) {
	node := list.head

	if printExtra {
		fmt.Println("Printing start")
	}

	if node == nil {
		fmt.Println("List does not exist")
	}

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}

	if printExtra {
		fmt.Println("Printing ends")
	}
}

func (list *LinkedList) Delete(value string) {
	node := list.head
	var start int = 0

	if node == nil {
		return
	}

	if node.value == value {
		list.head = node.next
		return
	}

	for node != nil {
		if node.next.value == value {
			node.next = node.next.next
			break
		}
		node = node.next
		start++
	}

}

func (list *LinkedList) ReverseLinkedList() {
	if list.head == nil {
		return
	}

	origHead := list.head
	headCopy := list.head
	nextNode := headCopy.next
	for headCopy != nil && nextNode != nil {
		if nextNode.next == nil {
			list.head = nextNode
		}

		next := nextNode.next
		nextNode.next = headCopy
		if headCopy == origHead {

			headCopy.next = nil

		}
		headCopy = nextNode
		nextNode = next
	}
}

func (list *LinkedList) CreateDummyLinkedList(data []string) *LinkedList {
	newList := &LinkedList{}
	for _, value := range data {
		newList.Add(value)
	}

	return newList

}

func helperReverse(node *Node) {
	if node == nil {
		return
	}

	if node.next == nil {
		fmt.Println(node.value)
		return
	}
	helperReverse(node.next)
	fmt.Println(node.value)
}

func (list *LinkedList) PrintReverseOrder() {

	helperReverse(list.head)

}

func (list *LinkedList) DeleteLinkedList() {
	list.head = nil
}
