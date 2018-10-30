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

// 头结点
type LinkedList struct {
	head *Node
}

// 新建节点
func NewNode(value ElementType) *Node {
	return &Node{value, nil}
}

// 新建链表
func NewLinkedList() *LinkedList {
	head := &Node{0, nil}
	return &LinkedList{head}
}

// 判断链表是否为空
func (list *LinkedList) IsEmpty() bool {
	return list.head.value == nil
}

// 链表长度
func (list *LinkedList) Length() int64 {
	return reflect.ValueOf(list.head.value).Int()
}

// 链表长度加1
func (list *LinkedList) SizeInc() {
	list.head.value = reflect.ValueOf(list.head.value).Int() + 1
}

// 链表长度减1
func (list *LinkedList) SizeDec() {
	list.head.value = reflect.ValueOf(list.head.value).Int() - 1
}

// 链表尾添加节点
func (list *LinkedList) Append(node *Node) {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	list.SizeInc()
}

// 链表头添加节点
func (list *LinkedList) Prepend(node *Node) {
	current := list.head
	node.next = current.next
	current.next = node
	list.SizeInc()
}

// 查询节点元素是否存在
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

// 删除节点元素（全部）
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

// 打印链表
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

// 打印节点信息
func (node *Node) PrintNode() {
	fmt.Printf("Node Value : %v\n", node.value)
}

// 寻找元素前节点
func (list *LinkedList) FindPreNode(value ElementType) (*Node, bool) {
	if list.IsEmpty() {
		fmt.Println("list is empty!")
		return nil, false
	}
	current := list.head
	for current.next != nil {
		if current.next.value == value {
			return current, true
		} else {
			current = current.next
		}
	}
	return nil, false
}

// 删除链表尾部第一个节点
func (list *LinkedList) Pop() (*Node, bool) {
	if list.IsEmpty() {
		fmt.Println("list is empty!")
		return nil, false
	}
	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	node := current.next
	current.next = nil
	list.SizeDec()
	return node, true
}

// 删除链表头第一个节点
func (list *LinkedList) Shift() (*Node, bool) {
	if list.IsEmpty() {
		fmt.Println("list is empty!")
		return nil, false
	}
	current := list.head.next
	list.head.next = current.next
	list.SizeDec()
	return current, true
}

// 单链表反转
func (list *LinkedList) Reverse() {
	if list.IsEmpty() || list.Length() == 1 {
		return
	}
	p := list.head.next
	q := p.next
	for q != nil {
		p.next = q.next

		q.next = list.head.next
		list.head.next = q

		q = p.next
	}
}
