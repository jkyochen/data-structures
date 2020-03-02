package array

import (
	"errors"
	"strconv"
)

// Array continuous data structure
type Array struct {
	data   []int
	length uint
}

// New new Array with capacity
func New(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (a *Array) isOutOfIndexRange(i uint) bool {
	if i >= uint(cap(a.data)) {
		return true
	}
	return false
}

// Get get element from index
func (a *Array) Get(i uint) (int, error) {
	if a.isOutOfIndexRange(i) {
		return 0, errors.New("Out of index range")
	}
	return a.data[i], nil
}

// Insert insert new element with index to Array
func (a *Array) Insert(i uint, ele int) error {

	if a.length == uint(cap(a.data)) {
		return errors.New("Array is full")
	}

	if a.isOutOfIndexRange(i) {
		return errors.New("Out of index range")
	}

	copy(a.data[i+1:], a.data[i:])
	a.data[i] = ele
	a.length++
	return nil
}

// InsertEnding insert element to arrray end
func (a *Array) InsertEnding(ele int) error {
	return a.Insert(a.length, ele)
}

// Delete delete element with index
func (a *Array) Delete(i uint) error {

	if a.isOutOfIndexRange(i) {
		return errors.New("Out of index range")
	}

	if a.length == 0 {
		return errors.New("Array is empty")
	}

	copy(a.data[i:], a.data[i+1:])
	a.length--
	return nil
}

// Len current Array length
func (a *Array) Len() uint {
	return a.length
}

func (a *Array) print() string {

	var str string
	for _, v := range a.data {
		str += strconv.Itoa(v) + "|"
	}
	return str
}
