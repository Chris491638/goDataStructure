package linkedList

import (
	"fmt"
	"reflect"
)

type ElementType interface{}

type Node struct {
	value ElementType
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(value ElementType) *Node {
	return &Node{value, nil}
}

func NewLinkedList() *LinkedList {
	head := &Node{0, nil}
	return &LinkedList{head}
}

func (list *LinkedList) IsEmpty() bool {
	return list.head.value == nil
}

func (list *LinkedList) Length() int64 {
	return reflect.ValueOf(list.head.value).Int()
}

func (list *LinkedList) SizeInc() {
	list.head.value = reflect.ValueOf(list.head.value).Int() + 1
}

func (list *LinkedList) SizeDec() {
	list.head.value = reflect.ValueOf(list.head.value).Int() - 1
}

func (list *LinkedList) Append(node *Node) {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	list.SizeInc()
}

func (list *LinkedList) Prepend(node *Node) {
	current := list.head
	node.next = current.next
	current.next = node
	list.SizeInc()
}

func (list *LinkedList) Find(value ElementType) (*Node, bool) {
	if list.IsEmpty() {
		fmt.Println("list is empty!")
		return nil, false
	}
	current := list.head
	for current.next != nil {
		if current.value == value {
			return current, true
		}
		current = current.next
	}
	if current.value == value {
		return current, true
	}
	return nil, false
}

func (list *LinkedList) Remove(value ElementType) {
	if list.IsEmpty() {
		fmt.Println("list is empty!")
	}
	current := list.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			list.SizeDec()
		} else {
			current = current.next
		}
	}

}

func (list *LinkedList) PrintList() {
	if list.IsEmpty() {
		fmt.Println("list is empty!")
		return
	}
	current := list.head.next
	fmt.Printf("Length : %v\n", list.head.value)
	index := 1
	for ; current.next != nil; index++ {
		fmt.Printf("index %d : %v\n", index, current.value)
		current = current.next
	}
	fmt.Printf("index %d : %v\n", index, current.value)
}

func (node *Node) PrintNode() {
	fmt.Printf("Node Value : %v\n", node.value)
}
