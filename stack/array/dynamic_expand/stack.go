package stack

import (
	"errors"
	"strconv"
)

const defaultCapacity = 3

// Stack last input first output LIFO data structure
type Stack struct {
	data   []int
	length int
}

// NewStack create new Stack
func NewStack() *Stack {
	return &Stack{
		data:   make([]int, defaultCapacity),
		length: 0,
	}
}

// Push add new element to Stack
func (s *Stack) Push(value int) error {
	if s.length == cap(s.data) {
		s.data = append(s.data, 0)
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
