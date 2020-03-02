package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	if NewStack() == nil {
		t.Fatal("NewStack did not work as expected.")
	}
}

func TestNewStackCap(t *testing.T) {
	if NewStackCap(3) == nil {
		t.Fatal("NewStackCap did not work as expected.")
	}
}

func TestPush(t *testing.T) {
	stack := NewStackCap(3)
	err := stack.Push(1)
	if err != nil {
		t.Fatal("Push did not work as expected.")
	}
	if stack.print() != "1|" {
		t.Fatal("Push did not work as expected.")
	}

	stack.Push(2)
	stack.Push(3)
	if stack.print() != "1|2|3|" {
		t.Fatal("Push did not work as expected.")
	}

	err = stack.Push(4)
	if err == nil || err.Error() != "Stack is full" {
		t.Fatal("Push did not work as expected.")
	}
}

func TestPop(t *testing.T) {
	stack := NewStackCap(3)
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
