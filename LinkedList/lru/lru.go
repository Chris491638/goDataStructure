package main

import (
	"goDataStructure/LinkedList/list"
)

func LRU(list *linkedList.LinkedList, value int, limit int64) {
	_, flag := list.Find(value)

	if flag {
		list.Remove(value)
	} else {
		if list.Length() >= limit {
			list.Pop()
		}
	}
	list.Prepend(linkedList.NewNode(value))
}

func main() {
	list := linkedList.NewLinkedList()
	for i := 1; i < 11; i++ {
		node := linkedList.NewNode(i)
		list.Prepend(node)
	}

	LRU(list, 20, 5)
	list.PrintList()

	LRU(list, 11, 11)
	list.PrintList()

	LRU(list, 13, 11)
	list.PrintList()

	LRU(list, 5, 20)
	list.PrintList()

}
