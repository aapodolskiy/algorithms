package main

import (
	"fmt"
	"strconv"
)

//////////////////////////////////////////////
////////// MyMinHeap implementation //////////

type MyMinHeapInterface interface {
	Empty() bool
	GetMin() int
	ExtractMin() int
	Insert(int)
}

type MyMinHeap struct {
	keys *[]int
	size int
}

func MyMinHeapConstructor() MyMinHeapInterface {
	size := 0
	keys := make([]int, size)
	return &MyMinHeap{&keys, size}
}

func (this *MyMinHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMinHeap) GetMin() int {
	return (*this.keys)[0]
}

func (this *MyMinHeap) ExtractMin() int {
	min := (*this.keys)[0]

	this.size--
	this._swap(0, this.size)
	this._minHeapify(0)

	return min
}

func (this *MyMinHeap) Insert(key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
	} else {
		*this.keys = append(*this.keys, key)
	}
	for i := this.size; i > 0 && (*this.keys)[i] < (*this.keys)[this._parent(i)]; i = this._parent(i) {
		this._swap(i, this._parent(i))
	}
	this.size++
}

func (this *MyMinHeap) _left(i int) int {
	return 2*(i+1) - 1
}
func (this *MyMinHeap) _right(i int) int {
	return 2*(i+1)
}
func (this *MyMinHeap) _parent(i int) int {
	return (i+1)/2 - 1
}
func (this *MyMinHeap) _swap(i, j int) {
	(*this.keys)[i], (*this.keys)[j] = (*this.keys)[j], (*this.keys)[i]
}

func (this *MyMinHeap) _minHeapify(i int) {
	l, r := this._left(i), this._right(i)
	smallest := i
	if l < this.size && (*this.keys)[l] < (*this.keys)[smallest] {
		smallest = l
	}
	if r < this.size && (*this.keys)[r] < (*this.keys)[smallest] {
		smallest = r
	}
	if smallest != i {
		this._swap(i, smallest)
		this._minHeapify(smallest)
	}
}

func (this *MyMinHeap) String() string {
	res := "minHeap:"
	for _, elem := range *this.keys {
		res += " " + strconv.Itoa(elem)
	}
	return res
}

////////// MyMinHeap implementation //////////
//////////////////////////////////////////////

func main() {
	h := MyMinHeapConstructor()
	h.Insert(4)
	h.Insert(1)
	h.Insert(5)
	fmt.Println(h.GetMin())
	fmt.Println(h.ExtractMin())
	h.Insert(7)
	h.Insert(7)
	h.Insert(17)
	fmt.Println(h.ExtractMin())
	h.Insert(10)
	fmt.Println(h)
}
