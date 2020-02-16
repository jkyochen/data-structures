package stack

import (
	"errors"
	"strconv"
)

const DEFAULT_CAPACITY = 10

type Stack struct {
	data   []int
	length int
}

func NewStack() *Stack {
	return &Stack{
		data:   make([]int, DEFAULT_CAPACITY, DEFAULT_CAPACITY),
		length: 0,
	}
}

func NewStackCap(capacity int) *Stack {
	return &Stack{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (s *Stack) Push(value int) error {
	if s.length == cap(s.data) {
		return errors.New("Stack is full")
	}

	s.data[s.length] = value
	s.length++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Stack is empty")
	}
	s.length--
	return s.data[s.length], nil
}

func (s *Stack) Print() string {

	var str string
	for i := 0; i < s.length; i++ {
		str += strconv.Itoa(s.data[i]) + "|"
	}
	return str
}
