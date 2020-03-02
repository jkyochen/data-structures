package stack

import (
	"errors"
	"strconv"
)

type Node struct {
	next *Node
	data int
}

type Stack struct {
	head   *Node
	length int
}

func NewStack() *Stack {
	return &Stack{
		head:   nil,
		length: 0,
	}
}

func (s *Stack) Push(value int) error {
	node := &Node{
		next: nil,
		data: value,
	}
	if s.head == nil {
		s.head = node
	} else {
		nextNode := s.head
		for nextNode.next != nil {
			nextNode = nextNode.next
		}
		nextNode.next = node
	}
	s.length++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}

	nextNode := s.head
	var preNode *Node
	for nextNode.next != nil {
		preNode = nextNode
		nextNode = nextNode.next
	}
	if preNode == nil {
		s.head = nil
	} else {
		preNode.next = nil
	}
	s.length--
	return nextNode.data, nil
}

func (s *Stack) Print() string {

	var str string
	nextNode := s.head
	for i := 0; i < s.length; i++ {
		str += strconv.Itoa(nextNode.data) + "|"
		nextNode = nextNode.next
	}
	return str
}
