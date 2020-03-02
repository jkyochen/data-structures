package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	if NewQueue() == nil {
		t.Fatal("NewQueue did not work as expected.")
	}
}

func TestNewQueueCap(t *testing.T) {
	if NewQueueCap(3) == nil {
		t.Fatal("NewQueueCap did not work as expected.")
	}
}

func TestEnqueue(t *testing.T) {
	queue := NewQueueCap(3)
	err := queue.Enqueue(1)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	if queue.Print() != "1|" {
		t.Fatal("Enqueue did not work as expected.")
	}

	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Print() != "1|2|3|" {
		t.Fatal("Enqueue did not work as expected.")
	}

	err = queue.Enqueue(4)
	if err == nil || err.Error() != "Queue is full" {
		t.Fatal("Enqueue did not work as expected.")
	}
}

func TestDequeue(t *testing.T) {
	queue := NewQueueCap(3)
	queue.Enqueue(1)
	v, err := queue.Dequeue()
	if err != nil || v != 1 {
		t.Fatal("Dequeue did not work as expected.")
	}

	_, err = queue.Dequeue()
	if err == nil || err.Error() != "Queue is empty" {
		t.Fatal("Enqueue did not work as expected.")
	}

	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	v, err = queue.Dequeue()
	if err != nil || v != 2 {
		t.Fatal("Dequeue did not work as expected.")
	}
	v, err = queue.Dequeue()
	if err != nil || v != 3 {
		t.Fatal("Dequeue did not work as expected.")
	}
}
