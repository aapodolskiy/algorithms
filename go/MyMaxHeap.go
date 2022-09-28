package main

import (
	"fmt"
)

//////////////////////////////////////////////
////////// MyMaxHeap implementation //////////

type MyHeapObject string //interface{}
var EmptyMyHeapObject MyHeapObject

type MyMaxHeap struct {
	objects *[]MyHeapObject
	keys    *[]int
	size    int
}

func (this *MyMaxHeap) Left(i int) int {
	return 2*(i+1) - 1
}

func (this *MyMaxHeap) Right(i int) int {
	return 2 * (i + 1)
}

func (this *MyMaxHeap) Parent(i int) int {
	return (i+1)/2 - 1
}

func swap[T any](a *[]T, i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (this *MyMaxHeap) maxHeapify(i int) {
	l, r := this.Left(i), this.Right(i)
	largest := i
	if l < this.size && (*this.keys)[largest] < (*this.keys)[l] {
		largest = l
	}
	if r < this.size && (*this.keys)[largest] < (*this.keys)[r] {
		largest = r
	}
	if largest != i {
		swap(this.keys, i, largest)
		swap(this.objects, i, largest)
		this.maxHeapify(largest)
	}
}

func MyMaxHeapConstructor() MyMaxHeap {
	size := 0
	keys := make([]int, size)
	objects := make([]MyHeapObject, size)
	return MyMaxHeap{&objects, &keys, size}
}

func (this *MyMaxHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMaxHeap) GetMaxObject() MyHeapObject {
	if this.Empty() {
		return EmptyMyHeapObject
	}
	return (*this.objects)[0]
}

func (this *MyMaxHeap) GetMaxKey() int {
	if this.Empty() {
		return -1
	}
	return (*this.keys)[0]
}

func (this *MyMaxHeap) ExtractMax() MyHeapObject {
	if this.Empty() {
		return EmptyMyHeapObject
	}

	objectWithMaxKey := (*this.objects)[0]

	this.size--
	swap(this.keys, 0, this.size)
	swap(this.objects, 0, this.size)
	this.maxHeapify(0)

	return objectWithMaxKey
}

func (this *MyMaxHeap) Insert(object MyHeapObject, key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
		(*this.objects)[this.size] = object
	} else {
		*this.keys = append(*this.keys, key)
		*this.objects = append(*this.objects, object)
	}
	for i := this.size; i > 0 && (*this.keys)[this.Parent(i)] < (*this.keys)[i]; i = this.Parent(i) {
		swap(this.keys, i, this.Parent(i))
		swap(this.objects, i, this.Parent(i))
	}
	this.size++
}

////////// MyMaxHeap implementation //////////
//////////////////////////////////////////////

func main() {
	h := MyMaxHeapConstructor()
	h.Insert("five", 5)
	fmt.Println(h.GetMaxObject(), h.GetMaxKey())
	h.ExtractMax()
	h.Insert("seven", 7)
	h.Insert("seventeen", 17)
	h.Insert("one more seventeen", 17)
	fmt.Println(h.ExtractMax())
	h.Insert("ten", 10)
	fmt.Println(h.keys, h.objects, h.size)
}
