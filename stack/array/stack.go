package stack

import (
	"errors"
	"strconv"
)

const defaultCapacity = 10

// Stack last input first output LIFO data structure
type Stack struct {
	data   []int
	length int
}

// NewStack create new Stack
func NewStack() *Stack {
	return &Stack{
		data:   make([]int, defaultCapacity, defaultCapacity),
		length: 0,
	}
}

// NewStackCap create new Stack with capacity
func NewStackCap(capacity int) *Stack {
	return &Stack{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// Push add new element to Stack
func (s *Stack) Push(value int) error {
	if s.length == cap(s.data) {
		return errors.New("Stack is full")
	}

	s.data[s.length] = value
	s.length++
	return nil
}

// Pop get element from Stack
func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}
	s.length--
	return s.data[s.length], nil
}

func (s *Stack) print() string {

	var str string
	for i := 0; i < s.length; i++ {
		str += strconv.Itoa(s.data[i]) + "|"
	}
	return str
}
