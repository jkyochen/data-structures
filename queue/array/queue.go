package queue

import (
	"errors"
	"strconv"
)

const DEFAULT_CAPACITY = 10

type Queue struct {
	data   []int
	length int
}

func NewQueue() *Queue {
	return &Queue{
		data:   make([]int, DEFAULT_CAPACITY, DEFAULT_CAPACITY),
		length: 0,
	}
}

func NewQueueCap(capacity int) *Queue {
	return &Queue{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (q *Queue) Enqueue(value int) error {
	if q.length == cap(q.data) {
		return errors.New("Queue is full")
	}

	q.data[q.length] = value
	q.length++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.length == 0 {
		return 0, errors.New("Queue is empty")
	}

	v := q.data[0]
	copy(q.data[0:], q.data[1:])
	q.length--
	return v, nil
}

func (q *Queue) Print() string {

	var str string
	for i := 0; i < q.length; i++ {
		str += strconv.Itoa(q.data[i]) + "|"
	}
	return str
}
