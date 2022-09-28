package main

import (
	"fmt"
)

//////////////////////////////////////////////
////////// MyMinHeap implementation //////////

type MyHeapObject interface{}

type MyMinHeap struct {
	objectsMap *map[int]MyHeapObject
	keys       *[]int
	size       int
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

func (this *MyMinHeap) minHeapify(i int) {
	l, r := this.Left(i), this.Right(i)
	smallest := i
	if l < this.size && (*this.keys)[l] < (*this.keys)[i] {
		smallest = l
	}
	if r < this.size && (*this.keys)[r] < (*this.keys)[i] {
		smallest = r
	}
	if smallest != i {
		(*this.keys)[i], (*this.keys)[smallest] = (*this.keys)[smallest], (*this.keys)[i]
		this.minHeapify(smallest)
	}
}

func MyMinHeapConstructor() MyMinHeap {
	size := 0
	keys := make([]int, size)
	objectsMap := make(map[int]MyHeapObject)
	return MyMinHeap{&objectsMap, &keys, size}
}

func (this *MyMinHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMinHeap) GetMin() (MyHeapObject, bool) {
	if this.Empty() {
		return nil, false
	}
	minKey := (*this.keys)[0]
	return (*this.objectsMap)[minKey], true
}

func (this *MyMinHeap) ExtractMin() (MyHeapObject, bool) {
	if this.Empty() {
		return nil, false
	}
	minKey := (*this.keys)[0]
	this.size--
	(*this.keys)[0] = (*this.keys)[this.size]
	this.minHeapify(0)
	objectWithMinKey := (*this.objectsMap)[minKey]
	delete(*this.objectsMap, minKey)
	return objectWithMinKey, true
}

func (this *MyMinHeap) Insert(value MyHeapObject, key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
	} else {
		*this.keys = append(*this.keys, key)
	}
	for i := this.size; i > 0 && (*this.keys)[i] < (*this.keys)[this.Parent(i)]; i = this.Parent(i) {
		(*this.keys)[i], (*this.keys)[this.Parent(i)] = (*this.keys)[this.Parent(i)], (*this.keys)[i]
	}
	this.size++
	(*this.objectsMap)[key] = value
}

////////// MyMinHeap implementation //////////
//////////////////////////////////////////////

func main() {
	h := MyMinHeapConstructor()
	h.Insert("five", 5)
	v, ok := h.GetMin()
	if ok {
		fmt.Println(v)
	}
	h.ExtractMin()
	h.Insert("seven", 7)
	h.Insert("seventeen", 17)
	h.Insert(-100500, 18)
	v, ok = h.ExtractMin()
	if ok {
		fmt.Println(v)
	}
	h.Insert("ten", 10)
	fmt.Println(h.keys, h.objectsMap, h.size)
}
