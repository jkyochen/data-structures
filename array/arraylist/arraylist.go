package arraylist

import (
	"errors"
	"strconv"
)

const defaultCapacity = 10

// ArrayList dynamic array
type ArrayList struct {
	data   []int
	length int
}

// NewArray create new arraylist
func NewArray() *ArrayList {
	return &ArrayList{
		data:   make([]int, defaultCapacity),
		length: 0,
	}
}

// NewArrayCap create new arraylist with capacity
func NewArrayCap(capacity int) *ArrayList {
	return &ArrayList{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (a *ArrayList) isOutOfIndexRange(i int) bool {
	return i >= a.length
}

// Get get element with index from arraylist
func (a *ArrayList) Get(i int) (int, error) {
	if a.isOutOfIndexRange(i) {
		return 0, errors.New("Out of index range")
	}
	return a.data[i], nil
}

// Insert insert new element to point index
func (a *ArrayList) Insert(i int, ele int) error {

	if a.length == cap(a.data) {
		a.data = append(a.data, 0)
	}

	copy(a.data[i+1:], a.data[i:])
	a.data[i] = ele
	a.length++
	return nil
}

// InsertEnding insert new element to arraylist end
func (a *ArrayList) InsertEnding(ele int) error {
	return a.Insert(a.length, ele)
}

// Delete delete element with index
func (a *ArrayList) Delete(i int) error {

	if a.length == 0 {
		return errors.New("ArrayList is empty")
	}

	if a.isOutOfIndexRange(i) {
		return errors.New("Out of index range")
	}

	copy(a.data[i:], a.data[i+1:])
	a.length--
	return nil
}

// Len current arraylist length
func (a *ArrayList) Len() int {
	return a.length
}

func (a *ArrayList) print() string {

	var str string
	for i := 0; i < a.length; i++ {
		str += strconv.Itoa(a.data[i]) + "|"
	}
	return str
}
