package main

import (
	"fmt"
	"strconv"
)

//////////////////////////////////////////////

type MySetInterface interface {
	IsEmpty() bool
	GetSize() int
	Insert(int)
	GetMax() int
	ExtractTop() int
}

type MyHeap struct {
	keys *[]int
	size int
	f    func(parent int, child int) bool
}

func _createMyHeap(f func(int, int) bool, input []int) *MyHeap {
	size := len(input)
	keys := make([]int, size)
	copy(keys, input)
	h := &MyHeap{&keys, size, f}
	for i := size / 2; i >= 0; i-- {
		h._heapify(i)
	}
	return h
}

func MyMaxHeapConstructor(input ...int) MyHeapInterface {
	return _createMyHeap(
		func(parent int, child int) bool { return parent > child },
		input,
	)
}

func MyMinHeapConstructor(input ...int) MyHeapInterface {
	return _createMyHeap(
		func(parent int, child int) bool { return parent < child },
		input,
	)
}

func (this *MyHeap) IsEmpty() bool {
	return this.size == 0
}

func (this *MyHeap) GetSize() int {
	return this.size
}

func (this *MyHeap) GetTop() int {
	return (*this.keys)[0]
}

func (this *MyHeap) ExtractTop() int {
	top := (*this.keys)[0]
	this.size--
	this._swap(0, this.size)
	this._heapify(0)
	return top
}

func (this *MyHeap) Insert(key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
	} else {
		*this.keys = append(*this.keys, key)
	}
	for i := this.size; i > 0 && !this.f((*this.keys)[this._parent(i)], (*this.keys)[i]); i = this._parent(i) {
		this._swap(i, this._parent(i))
	}
	this.size++
}

func (this *MyHeap) _left(i int) int {
	return 2*(i+1) - 1
}
func (this *MyHeap) _right(i int) int {
	return 2 * (i + 1)
}
func (this *MyHeap) _parent(i int) int {
	return (i+1)/2 - 1
}
func (this *MyHeap) _swap(i, j int) {
	(*this.keys)[i], (*this.keys)[j] = (*this.keys)[j], (*this.keys)[i]
}
func (this *MyHeap) _heapify(i int) {
	l, r := this._left(i), this._right(i)
	parent := i
	if l < this.size && !this.f((*this.keys)[parent], (*this.keys)[l]) {
		parent = l
	}
	if r < this.size && !this.f((*this.keys)[parent], (*this.keys)[r]) {
		parent = r
	}
	if parent != i {
		this._swap(i, parent)
		this._heapify(parent)
	}
}

func (this *MyHeap) String() string {
	res := "heap:"
	for i, elem := range *this.keys {
		if i >= this.size {
			break
		}
		res += " " + strconv.Itoa(elem)
	}
	return res
}

////////// MyHeap implementation
//////////////////////////////////////////////

func main() {
	maxHeap := MyMaxHeapConstructor([]int{4, 7, 1, 6}...)
	fmt.Println(maxHeap.GetSize(), maxHeap.IsEmpty(), maxHeap.GetTop())
	maxHeap.Insert(5)
	fmt.Println(maxHeap)
	maxHeap.Insert(7)
	fmt.Println(maxHeap.ExtractTop())
	fmt.Println(maxHeap)
	for !maxHeap.IsEmpty() {
		fmt.Print(maxHeap.ExtractTop(), " ")
	}

	fmt.Println("\n\n")

	minHeap := MyMinHeapConstructor()
	fmt.Println(minHeap.GetSize(), minHeap.IsEmpty())
	minHeap.Insert(5)
	minHeap.Insert(0)
	minHeap.Insert(2)
	minHeap.Insert(4)
	fmt.Println(minHeap)
	minHeap.Insert(7)
	fmt.Println(minHeap.ExtractTop())
	fmt.Println(minHeap)
	for !minHeap.IsEmpty() {
		fmt.Print(minHeap.ExtractTop(), " ")
	}
}
