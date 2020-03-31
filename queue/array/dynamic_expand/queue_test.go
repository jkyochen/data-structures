package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	if New() == nil {
		t.Fatal("New did not work as expected.")
	}
}

func TestEnqueue(t *testing.T) {
	queue := New()
	err := queue.Enqueue(1)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	if queue.print() != "1|" {
		t.Fatal("Enqueue did not work as expected.")
	}

	err = queue.Enqueue(2)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	err = queue.Enqueue(3)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	err = queue.Enqueue(4)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}

	if queue.print() != "1|2|3|4|" {
		t.Fatal("Enqueue did not work as expected.")
	}
}

func TestDequeue(t *testing.T) {
	queue := New()
	err := queue.Enqueue(1)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}

	v, err := queue.Dequeue()
	if err != nil || v != 1 {
		t.Fatal("Dequeue did not work as expected.")
	}

	_, err = queue.Dequeue()
	if err == nil || err.Error() != "Queue is empty" {
		t.Fatal("Enqueue did not work as expected.")
	}

	err = queue.Enqueue(2)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	err = queue.Enqueue(3)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	v, err = queue.Dequeue()
	if err != nil || v != 2 {
		t.Fatal("Dequeue did not work as expected.")
	}

	err = queue.Enqueue(4)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}

	v, err = queue.Dequeue()
	if err != nil || v != 3 {
		t.Fatal("Dequeue did not work as expected.")
	}
}

func TestLen(t *testing.T) {
	queue := New()
	err := queue.Enqueue(1)
	if err != nil {
		t.Fatal("Enqueue did not work as expected.")
	}
	if queue.Len() != 1 {
		t.Fatal("Len did not work as expected.")
	}

	v, err := queue.Dequeue()
	if err != nil || v != 1 {
		t.Fatal("Dequeue did not work as expected.")
	}
	if queue.Len() != 0 {
		t.Fatal("Len did not work as expected.")
	}
}
