package queue

import (
	"errors"
	"strconv"
)

type Node struct {
	next *Node
	data int
}

type Queue struct {
	head   *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		head:   nil,
		length: 0,
	}
}

func (q *Queue) Enqueue(value int) error {
	node := &Node{
		next: nil,
		data: value,
	}
	if q.head == nil {
		q.head = node
	} else {
		nextNode := q.head
		for nextNode.next != nil {
			nextNode = nextNode.next
		}
		nextNode.next = node
	}
	q.length++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.length == 0 {
		return 0, errors.New("Queue is empty")
	}

	v := q.head.data
	q.head = q.head.next
	q.length--
	return v, nil
}

func (q *Queue) Print() string {

	var str string
	nextNode := q.head
	for i := 0; i < q.length; i++ {
		str += strconv.Itoa(nextNode.data) + "|"
		nextNode = nextNode.next
	}
	return str
}
