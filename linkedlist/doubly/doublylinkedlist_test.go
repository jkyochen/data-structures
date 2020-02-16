package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	if list == nil {
		t.Fatal("New did not work as expected.")
	}
	if list.Print() != "" {
		t.Fatal("New did not work as expected.")
	}
}

func TestAdd(t *testing.T) {
	list := New()
	list.Add(&Node{
		data: 1,
	})
	if list.Print() != "1|1|" {
		t.Fatal("Add did not work as expected.")
	}
	list.Add(&Node{
		data: 2,
	})
	if list.Print() != "1|2|2|1|" {
		t.Fatal("Add did not work as expected.")
	}
}

func TestInsert(t *testing.T) {
	list := New()
	list.Add(&Node{
		data: 1,
	})
	err := list.Insert(1, 10)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	if list.Print() != "1|10|10|1|" {
		t.Fatal("Insert did not work as expected.")
	}

	err = list.Insert(1, 2)
	if err != nil {
		t.Fatal("Insert did not work as expected.")
	}
	if list.Print() != "1|2|10|10|2|1|" {
		t.Fatal("Insert did not work as expected.")
	}

	err = list.Insert(11, 3)
	if err == nil || err.Error() != "don't find this value" {
		t.Fatal("Insert did not work as expected.")
	}
}

func TestDelete(t *testing.T) {
	list := New()
	list.Add(&Node{
		data: 1,
	})
	err := list.Delete(1)
	if err != nil {
		t.Fatal("Delete did not work as expected.")
	}
	list.Add(&Node{
		data: 1,
	})
	list.Add(&Node{
		data: 2,
	})
	list.Add(&Node{
		data: 3,
	})
	err = list.Delete(2)
	if err != nil {
		t.Fatal("Delete did not work as expected.")
	}
	if list.Print() != "1|3|3|1|" {
		t.Fatal("Delete did not work as expected.")
	}

	err = list.Delete(3)
	if err != nil {
		t.Fatal("Delete did not work as expected.")
	}
	if list.Print() != "1|1|" {
		t.Fatal("Delete did not work as expected.")
	}

	err = list.Delete(10)
	if err == nil || err.Error() != "don't find this value" {
		t.Fatal("Delete did not work as expected.")
	}
}

func TestSize(t *testing.T) {
	list := New()
	list.Add(&Node{
		data: 1,
	})
	if list.Len() != 1 {
		t.Fatal("Len did not work as expected.")
	}

	list.Insert(1, 2)
	if list.Len() != 2 {
		t.Fatal("Len did not work as expected.")
	}

	list.Delete(1)
	if list.Len() != 1 {
		t.Fatal("Len did not work as expected.")
	}
}
