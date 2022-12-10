package main

import (
	"fmt"
	"strconv"
)

//////////////////////////////////////////////
////////// MyMaxHeap implementation //////////

type MyMaxHeapInterface interface {
	Empty() bool
	GetMax() int
	ExtractMax() int
	Insert(int)
}

type MyMaxHeap struct {
	keys *[]int
	size int
}

func MyMaxHeapConstructor() MyMaxHeapInterface {
	size := 0
	keys := make([]int, size)
	return &MyMaxHeap{&keys, size}
}

func (this *MyMaxHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMaxHeap) GetMax() int {
	return (*this.keys)[0]
}

func (this *MyMaxHeap) ExtractMax() int {
	max := (*this.keys)[0]

	this.size--
	this._swap(0, this.size)
	this._maxHeapify(0)

	return max
}

func (this *MyMaxHeap) Insert(key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
	} else {
		*this.keys = append(*this.keys, key)
	}
	for i := this.size; i > 0 && (*this.keys)[this._parent(i)] < (*this.keys)[i]; i = this._parent(i) {
		this._swap(i, this._parent(i))
	}
	this.size++
}

func (this *MyMaxHeap) _left(i int) int {
	return 2*(i+1) - 1
}
func (this *MyMaxHeap) _right(i int) int {
	return 2*(i+1)
}
func (this *MyMaxHeap) _parent(i int) int {
	return (i+1)/2 - 1
}
func (this *MyMaxHeap) _swap(i, j int) {
	(*this.keys)[i], (*this.keys)[j] = (*this.keys)[j], (*this.keys)[i]
}

func (this *MyMaxHeap) _maxHeapify(i int) {
	l, r := this._left(i), this._right(i)
	largest := i
	if l < this.size && (*this.keys)[largest] < (*this.keys)[l] {
		largest = l
	}
	if r < this.size && (*this.keys)[largest] < (*this.keys)[r] {
		largest = r
	}
	if largest != i {
		this._swap(i, largest)
		this._maxHeapify(largest)
	}
}

func (this *MyMaxHeap) String() string {
	res := "maxHeap:"
	for _, elem := range *this.keys {
		res += " " + strconv.Itoa(elem)
	}
	return res
}

////////// MyMaxHeap implementation //////////
//////////////////////////////////////////////

func main() {
	h := MyMaxHeapConstructor()
	h.Insert(5)
	fmt.Println(h.GetMax())
	h.Insert(7)
	fmt.Println(h.ExtractMax())
	h.Insert(27)
	h.Insert(17)
	h.Insert(17)
	fmt.Println(h.ExtractMax())
	h.Insert(10)
	fmt.Println(h)
}
