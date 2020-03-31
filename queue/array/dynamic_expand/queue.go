package queue

import (
	"errors"
	"strconv"
)

const defaultCapacity = 3

// Queue last input first output LIFO data structure
type Queue struct {
	data []int
}

// New create new Queue
func New() *Queue {
	return &Queue{
		data: make([]int, 0, defaultCapacity),
	}
}

// Enqueue add new element to Queue
func (s *Queue) Enqueue(value int) error {
	s.data = append(s.data, value)
	return nil
}

// Dequeue get element from Queue
func (s *Queue) Dequeue() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Queue is empty")
	}
	v := s.data[0]
	s.data = s.data[1:]
	return v, nil
}

// Len return Queue length
func (s *Queue) Len() int {
	return len(s.data)
}

func (s *Queue) print() string {

	var str string
	for i := 0; i < len(s.data); i++ {
		str += strconv.Itoa(s.data[i]) + "|"
	}
	return str
}
