package linkedlist

import (
	"errors"
	"strconv"
)

type node struct {
	next *node
	data int
}

// SinglyLinkedList singly linked list
type SinglyLinkedList struct {
	head   *node
	length uint
}

// New new SinglyLinkedList
func New() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   nil,
		length: 0,
	}
}

// Add new node
func (list *SinglyLinkedList) Add(node *node) {
	if list.head == nil {
		list.head = node
	} else {
		curNode := list.head
		for curNode.next != nil {
			curNode = curNode.next
		}
		curNode.next = node
	}
	list.length++
}

// Insert insert new element, when equal give value
func (list *SinglyLinkedList) Insert(value int, newValue int) error {
	curNode := list.head
	for curNode != nil && curNode.data != value {
		curNode = curNode.next
	}
	if curNode == nil {
		return errors.New("don't find this value")
	}
	node := &node{
		next: curNode.next,
		data: newValue,
	}
	curNode.next = node
	list.length++
	return nil
}

// Delete delete give value
func (list *SinglyLinkedList) Delete(value int) error {
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
	}
	list.length--
	return nil
}

// Len current SinglyLinkedList length
func (list *SinglyLinkedList) Len() uint {
	return list.length
}

func (list *SinglyLinkedList) print() string {
	var str string
	curNode := list.head
	for curNode != nil {
		str += strconv.Itoa(curNode.data) + "|"
		curNode = curNode.next
	}
	return str
}
