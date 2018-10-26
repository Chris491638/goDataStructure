package main

import (
	"goDataStructure/LinkedList/list"
)

func main() {
	p := linkedList.NewLinkedList()

	node := linkedList.NewNode(1)
	p.Append(node)
	p.PrintList()

	node = linkedList.NewNode("str")
	p.Prepend(node)
	p.PrintList()

	node = linkedList.NewNode(2)
	p.Append(node)
	p.PrintList()

	fNode, flag := p.Find("str")
	if flag {
		fNode.PrintNode()
	}

	p.Remove("str")
	p.PrintList()
}
