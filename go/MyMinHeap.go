package main

import (
	"fmt"
)

//////////////////////////////////////////////
////////// MyMinHeap implementation //////////

type MyHeapObject string     //interface{}
const EmptyMyHeapObject = "" //0

type MyMinHeap struct {
	objects *[]MyHeapObject
	keys    *[]int
	size    int
}

func (this *MyMinHeap) Left(i int) int {
	return 2*(i+1) - 1
}

func (this *MyMinHeap) Right(i int) int {
	return 2 * (i + 1)
}

func (this *MyMinHeap) Parent(i int) int {
	return (i+1)/2 - 1
}

func swap[T any](a *[]T, i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (this *MyMinHeap) minHeapify(i int) {
	l, r := this.Left(i), this.Right(i)
	smallest := i
	if l < this.size && (*this.keys)[l] < (*this.keys)[smallest] {
		smallest = l
	}
	if r < this.size && (*this.keys)[r] < (*this.keys)[smallest] {
		smallest = r
	}
	if smallest != i {
		swap(this.keys, i, smallest)
		swap(this.objects, i, smallest)
		this.minHeapify(smallest)
	}
}

func MyMinHeapConstructor() MyMinHeap {
	size := 0
	keys := make([]int, size)
	objects := make([]MyHeapObject, size)
	return MyMinHeap{&objects, &keys, size}
}

func (this *MyMinHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMinHeap) GetMinObject() MyHeapObject {
	if this.Empty() {
		return EmptyMyHeapObject
	}
	return (*this.objects)[0]
}

func (this *MyMinHeap) GetMinKey() int {
	if this.Empty() {
		return -1
	}
	return (*this.keys)[0]
}

func (this *MyMinHeap) ExtractMin() MyHeapObject {
	if this.Empty() {
		return EmptyMyHeapObject
	}

	objectWithMinKey := (*this.objects)[0]

	this.size--
	swap(this.keys, 0, this.size)
	swap(this.objects, 0, this.size)
	this.minHeapify(0)

	return objectWithMinKey
}

func (this *MyMinHeap) Insert(object MyHeapObject, key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
		(*this.objects)[this.size] = object
	} else {
		*this.keys = append(*this.keys, key)
		*this.objects = append(*this.objects, object)
	}
	for i := this.size; i > 0 && (*this.keys)[i] < (*this.keys)[this.Parent(i)]; i = this.Parent(i) {
		swap(this.keys, i, this.Parent(i))
		swap(this.objects, i, this.Parent(i))
	}
	this.size++
}

////////// MyMinHeap implementation //////////
//////////////////////////////////////////////

func main() {
	h := MyMinHeapConstructor()
	h.Insert("five", 5)
	fmt.Println(h.GetMinObject(), h.GetMinKey())
	h.ExtractMin()
	h.Insert("seven", 7)
	h.Insert("one more seven", 7)
	h.Insert("seventeen", 17)
	fmt.Println(h.ExtractMin())
	h.Insert("ten", 10)
	fmt.Println(h.keys, h.objects, h.size)
}
