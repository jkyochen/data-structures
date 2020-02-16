package linkedlist

import (
	"errors"
	"strconv"
)

type Node struct {
	before *Node
	next   *Node
	data   int
}

type DoublyLinkedList struct {
	head   *Node
	length uint
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		length: 0,
	}
}

func (list *DoublyLinkedList) Add(node *Node) {
	if list.head == nil {
		list.head = node
	} else {
		curNode := list.head
		for curNode.next != nil {
			curNode = curNode.next
		}
		node.before = curNode
		curNode.next = node
	}
	list.length++
}

func (list *DoublyLinkedList) Insert(value int, newValue int) error {
	curNode := list.head
	for curNode != nil && curNode.data != value {
		curNode = curNode.next
	}
	if curNode == nil {
		return errors.New("don't find this value")
	}

	node := &Node{
		before: curNode,
		next:   curNode.next,
		data:   newValue,
	}
	if curNode.next != nil {
		curNode.next.before = node
	}
	curNode.next = node
	list.length++
	return nil
}

func (list *DoublyLinkedList) Delete(value int) error {
	curNode := list.head
	var befNode *Node
	for curNode != nil && curNode.data != value {
		befNode = curNode
		curNode = curNode.next
	}
	if curNode == nil {
		return errors.New("don't find this value")
	}
	if befNode == nil {
		list.head = nil
	} else {
		befNode.next = curNode.next
		if curNode.next != nil {
			curNode.next.before = befNode
		}
	}
	list.length--
	return nil
}

func (list *DoublyLinkedList) Len() uint {
	return list.length
}

func (list *DoublyLinkedList) Print() string {
	var str string
	var lastNode *Node
	curNode := list.head
	for curNode != nil {
		str += strconv.Itoa(curNode.data) + "|"
		lastNode = curNode
		curNode = curNode.next
	}

	curNode = lastNode
	for curNode != nil {
		str += strconv.Itoa(curNode.data) + "|"
		curNode = curNode.before
	}
	return str
}
