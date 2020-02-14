package arraylist

import (
	"errors"
	"strconv"
)

const DEFAULT_CAPACITY = 10

type ArrayList struct {
	data   []int
	length uint
}

func NewArray() *ArrayList {
	return &ArrayList{
		data:   make([]int, DEFAULT_CAPACITY, DEFAULT_CAPACITY),
		length: 0,
	}
}

func NewArrayCap(capacity uint) *ArrayList {
	return &ArrayList{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (a *ArrayList) isOutOfIndexRange(i uint) bool {
	if i >= uint(cap(a.data)) {
		return true
	}
	return false
}

func (a *ArrayList) Get(i uint) (int, error) {
	if a.isOutOfIndexRange(i) {
		return 0, errors.New("Out of index range")
	}
	return a.data[i], nil
}

func (a *ArrayList) Insert(i uint, ele int) error {

	if a.length == uint(cap(a.data)) {
		a.data = append(a.data, 0)
	}

	if a.isOutOfIndexRange(i) {
		return errors.New("Out of index range")
	}

	copy(a.data[i+1:], a.data[i:])
	a.data[i] = ele
	a.length++
	return nil
}

func (a *ArrayList) InsertEnding(ele int) error {
	return a.Insert(a.length, ele)
}

func (a *ArrayList) Delete(i uint) error {

	if a.isOutOfIndexRange(i) {
		return errors.New("Out of index range")
	}

	if a.length == 0 {
		return errors.New("ArrayList is empty")
	}

	copy(a.data[i:], a.data[i+1:])
	a.length--
	return nil
}

func (a *ArrayList) Len() uint {
	return a.length
}

func (a *ArrayList) Print() string {

	var str string
	for _, v := range a.data {
		str += strconv.Itoa(v) + "|"
	}
	return str
}
