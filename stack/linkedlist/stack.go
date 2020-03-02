package stack

import (
	"errors"
	"strconv"
)

type node struct {
	next *node
	data int
}

// Stack last input first output LIFO data structure
type Stack struct {
	head   *node
	length int
}

// NewStack create new Stack
func NewStack() *Stack {
	return &Stack{
		head:   nil,
		length: 0,
	}
}

// Push add new element to Stack
func (s *Stack) Push(value int) error {
	newNode := &node{
		next: nil,
		data: value,
	}
	if s.head == nil {
		s.head = newNode
	} else {
		nextNode := s.head
		for nextNode.next != nil {
			nextNode = nextNode.next
		}
		nextNode.next = newNode
	}
	s.length++
	return nil
}

// Pop get element from Stack
func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}

	nextNode := s.head
	var preNode *node
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

func (s *Stack) print() string {

	var str string
	nextNode := s.head
	for i := 0; i < s.length; i++ {
		str += strconv.Itoa(nextNode.data) + "|"
		nextNode = nextNode.next
	}
	return str
}
