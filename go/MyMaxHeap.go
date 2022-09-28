package main

import (
	"fmt"
)

//////////////////////////////////////////////
////////// MyMaxHeap implementation //////////

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

func (this *MyMinHeap) maxHeapify(i int) {
	l, r := this.Left(i), this.Right(i)
	largest := i
	if l < this.size && (*this.keys)[i] < (*this.keys)[l] {
		largest = l
	}
	if r < this.size && (*this.keys)[largest] < (*this.keys)[r] {
		largest = r
	}
	if largest != i {
		(*this.keys)[i], (*this.keys)[largest] = (*this.keys)[largest], (*this.keys)[i]
		this.maxHeapify(largest)
	}
}

func MyMaxHeapConstructor() MyMinHeap {
	size := 0
	keys := make([]int, size)
	objectsMap := make(map[int]MyHeapObject)
	return MyMinHeap{&objectsMap, &keys, size}
}

func (this *MyMinHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMinHeap) GetMax() (MyHeapObject, bool) {
	if this.Empty() {
		return nil, false
	}
	maxKey := (*this.keys)[0]
	return (*this.objectsMap)[maxKey], true
}

func (this *MyMinHeap) ExtractMax() (MyHeapObject, bool) {
	if this.Empty() {
		return nil, false
	}
	maxKey := (*this.keys)[0]
	this.size--
	(*this.keys)[0] = (*this.keys)[this.size]
	this.maxHeapify(0)
	objectWithMaxKey := (*this.objectsMap)[maxKey]
	delete(*this.objectsMap, maxKey)
	return objectWithMaxKey, true
}

func (this *MyMinHeap) Insert(value MyHeapObject, key int) {
	if len(*this.keys) > this.size {
		(*this.keys)[this.size] = key
	} else {
		*this.keys = append(*this.keys, key)
	}
	for i := this.size; i > 0 && (*this.keys)[this.Parent(i)] < (*this.keys)[i]; i = this.Parent(i) {
		(*this.keys)[i], (*this.keys)[this.Parent(i)] = (*this.keys)[this.Parent(i)], (*this.keys)[i]
	}
	this.size++
	(*this.objectsMap)[key] = value
}

////////// MyMaxHeap implementation //////////
//////////////////////////////////////////////

func main() {
	h := MyMaxHeapConstructor()
	h.Insert("five", 5)
	v, ok := h.GetMax()
	if ok {
		fmt.Println(v)
	}
	h.ExtractMax()
	h.Insert("seven", 7)
	h.Insert("seventeen", 17)
	h.Insert(-100500, 18)
	v, ok = h.ExtractMax()
	if ok {
		fmt.Println(v)
	}
	h.Insert("ten", 10)
	fmt.Println(h.keys, h.objectsMap, h.size)
}
