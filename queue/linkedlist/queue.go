package queue

import (
	"errors"
	"strconv"
)

type node struct {
	next *node
	data int
}

// Queue first input first output FIFO data structure
type Queue struct {
	head   *node
	length int
}

// NewQueue create new Queue
func NewQueue() *Queue {
	return &Queue{
		head:   nil,
		length: 0,
	}
}

// Enqueue add new element to Queue
func (q *Queue) Enqueue(value int) error {
	newNode := &node{
		next: nil,
		data: value,
	}
	if q.head == nil {
		q.head = newNode
	} else {
		nextNode := q.head
		for nextNode.next != nil {
			nextNode = nextNode.next
		}
		nextNode.next = newNode
	}
	q.length++
	return nil
}

// Dequeue get element from Queue
func (q *Queue) Dequeue() (int, error) {
	if q.length == 0 {
		return 0, errors.New("Queue is empty")
	}

	v := q.head.data
	q.head = q.head.next
	q.length--
	return v, nil
}

func (q *Queue) print() string {

	var str string
	nextNode := q.head
	for i := 0; i < q.length; i++ {
		str += strconv.Itoa(nextNode.data) + "|"
		nextNode = nextNode.next
	}
	return str
}
