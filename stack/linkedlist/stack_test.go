package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := NewStack()
	err := stack.Push(1)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}
	if stack.Print() != "1|" {
		t.Fatal("Push did not work as expected.")
	}

	stack.Push(2)
	stack.Push(3)
	if stack.Print() != "1|2|3|" {
		t.Fatal("Push did not work as expected.")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	v, err := stack.Pop()
	if err != nil || v != 1 {
		t.Fatal("Pop did not work as expected.")
	}

	_, err = stack.Pop()
	if err == nil || err.Error() != "Stack is empty" {
		t.Fatal("Push did not work as expected.")
	}

	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	v, err = stack.Pop()
	if err != nil || v != 4 {
		t.Fatal("Pop did not work as expected.")
	}
}
