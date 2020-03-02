package linkedlist

import (
	"errors"
	"strconv"
)

type node struct {
	next *node
	data int
}

// CircularLinkedList circular linked list
type CircularLinkedList struct {
	head   *node
	length int
}

// New new CircularLinkedList
func New() *CircularLinkedList {
	return &CircularLinkedList{
		head:   nil,
		length: 0,
	}
}

// Add new node
func (list *CircularLinkedList) Add(node *node) {
	if list.head == nil {
		list.head = node
		list.head.next = list.head
	} else {
		curNode := list.head
		for i := 1; i <= int(list.length-1); i++ {
			curNode = curNode.next
		}
		curNode.next = node
	}
	list.length++
}

// Insert insert new element, when equal give value
func (list *CircularLinkedList) Insert(value int, newValue int) error {
	curNode := list.head
	i := 1
	for ; i <= list.length && curNode.data != value; i++ {
		curNode = curNode.next
	}
	if i > list.length {
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
func (list *CircularLinkedList) Delete(value int) error {
	curNode := list.head
	var befNode *node
	i := 1
	for ; i <= list.length && curNode.data != value; i++ {
		befNode = curNode
		curNode = curNode.next
	}
	if i > list.length {
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

// Len current CircularLinkedList length
func (list *CircularLinkedList) Len() int {
	return list.length
}

func (list *CircularLinkedList) print() string {
	var str string
	curNode := list.head
	for i := 1; i <= list.length; i++ {
		str += strconv.Itoa(curNode.data) + "|"
		curNode = curNode.next
	}
	return str
}
