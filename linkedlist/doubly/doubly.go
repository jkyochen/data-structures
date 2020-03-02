package linkedlist

import (
	"errors"
	"strconv"
)

type node struct {
	before *node
	next   *node
	data   int
}

// DoublyLinkedList doubly linked list
type DoublyLinkedList struct {
	head   *node
	length uint
}

// New new DoublyLinkedList
func New() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		length: 0,
	}
}

// Add new node
func (list *DoublyLinkedList) Add(node *node) {
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

// Insert insert new element, when equal give value
func (list *DoublyLinkedList) Insert(value int, newValue int) error {
	curNode := list.head
	for curNode != nil && curNode.data != value {
		curNode = curNode.next
	}
	if curNode == nil {
		return errors.New("don't find this value")
	}

	node := &node{
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

// Delete delete give value
func (list *DoublyLinkedList) Delete(value int) error {
	curNode := list.head
	var befNode *node
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

// Len current DoublyLinkedList length
func (list *DoublyLinkedList) Len() uint {
	return list.length
}

func (list *DoublyLinkedList) print() string {
	var str string
	var lastNode *node
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
