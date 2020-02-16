package linkedlist

import (
	"errors"
	"strconv"
)

type Node struct {
	next *Node
	data int
}

type SinglyLinkedList struct {
	head   *Node
	length uint
}

func New() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   nil,
		length: 0,
	}
}

func (list *SinglyLinkedList) Add(node *Node) {
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

func (list *SinglyLinkedList) Insert(value int, newValue int) error {
	curNode := list.head
	for curNode != nil && curNode.data != value {
		curNode = curNode.next
	}
	if curNode == nil {
		return errors.New("don't find this value")
	}
	node := &Node{
		next: curNode.next,
		data: newValue,
	}
	curNode.next = node
	list.length++
	return nil
}

func (list *SinglyLinkedList) Delete(value int) error {
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
	}
	list.length--
	return nil
}

func (list *SinglyLinkedList) Len() uint {
	return list.length
}

func (list *SinglyLinkedList) Print() string {
	var str string
	curNode := list.head
	for curNode != nil {
		str += strconv.Itoa(curNode.data) + "|"
		curNode = curNode.next
	}
	return str
}
