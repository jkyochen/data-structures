package queue

import (
	"errors"
	"strconv"
)

const defaultCapacity = 10

// Queue first input first output FIFO data structure
type Queue struct {
	data   []int
	length int
}

// NewQueue create new Queue
func NewQueue() *Queue {
	return &Queue{
		data:   make([]int, defaultCapacity, defaultCapacity),
		length: 0,
	}
}

// NewQueueCap create new Queue with capacity
func NewQueueCap(capacity int) *Queue {
	return &Queue{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// Enqueue add new element to Queue
func (q *Queue) Enqueue(value int) error {
	if q.length == cap(q.data) {
		return errors.New("Queue is full")
	}

	q.data[q.length] = value
	q.length++
	return nil
}

// Dequeue get element from Queue
func (q *Queue) Dequeue() (int, error) {
	if q.length == 0 {
		return 0, errors.New("Queue is empty")
	}

	v := q.data[0]
	copy(q.data[0:], q.data[1:])
	q.length--
	return v, nil
}

func (q *Queue) print() string {

	var str string
	for i := 0; i < q.length; i++ {
		str += strconv.Itoa(q.data[i]) + "|"
	}
	return str
}
