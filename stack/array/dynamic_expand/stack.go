package stack

import (
	"errors"
	"strconv"
)

const defaultCapacity = 3

// Stack last input first output LIFO data structure
type Stack struct {
	data []int
}

// New create new Stack
func New() *Stack {
	return &Stack{
		data: make([]int, 0, defaultCapacity),
	}
}

// Push add new element to Stack
func (s *Stack) Push(value int) error {
	s.data = append(s.data, value)
	return nil
}

// Pop get element from Stack
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

// Len return stack length
func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) print() string {

	var str string
	for i := 0; i < len(s.data); i++ {
		str += strconv.Itoa(s.data[i]) + "|"
	}
	return str
}
