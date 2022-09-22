package main

import (
	"fmt"
)

//////////////////////////////////////////////
////////// MyMaxHeap implementation //////////

type MyMaxHeap struct {
	data	*[]int
	size	int
}

func (this *MyMaxHeap) Left(i int) int {
	return 2*(i+1)-1
}

func (this *MyMaxHeap) Right(i int) int {
	return 2*(i+1)
}

func (this *MyMaxHeap) Parent(i int) int {
	return (i+1)/2-1
}

func (this *MyMaxHeap) maxHeapify(i int) {
	l, r := this.Left(i), this.Right(i)
	largest := i
	if l < this.size && (*this.data)[i] < (*this.data)[l] {
		largest = l
	}
	if r < this.size && (*this.data)[largest] < (*this.data)[r] {
		largest = r
	}
	if largest != i {
		(*this.data)[i], (*this.data)[largest] = (*this.data)[largest], (*this.data)[i]
		this.maxHeapify(largest)
	}
}

func MyMaxHeapConstructor(a *[]int) MyMaxHeap {
	size := len(*a)
	data := make([]int, size)
	copy(data, *a)
	heap := MyMaxHeap{&data, size}
	for i := heap.size/2+1; i >= 0; i-- {
		heap.maxHeapify(i)
	}
    return heap
}

func (this *MyMaxHeap) Empty() bool {
	return this.size == 0
}

func (this *MyMaxHeap) GetMax() int {
	return (*this.data)[0]
}

func (this *MyMaxHeap) ExtractMax() int {
	if this.Empty() {
		return -1
	}
	max := (*this.data)[0]
	this.size--
	(*this.data)[0] = (*this.data)[this.size]
	this.maxHeapify(0)
	return max
}

func (this *MyMaxHeap) Insert(x int) {
	if len(*this.data) > this.size {
		(*this.data)[this.size] = x
	} else {
		*this.data = append(*this.data, x)
	}
	for i := this.size; i > 0 && (*this.data)[this.Parent(i)] < (*this.data)[i]; i = this.Parent(i) {
		(*this.data)[i], (*this.data)[this.Parent(i)] = (*this.data)[this.Parent(i)], (*this.data)[i]
	}
	this.size++
}

////////// MyMaxHeap implementation //////////
//////////////////////////////////////////////

func main() {
	a := []int{1,2,3,4,5,6,9,10}
	h := MyMaxHeapConstructor(&a)
	h.Insert(16)
	fmt.Println(h.GetMax())
	h.ExtractMax()
	h.Insert(7)
	h.ExtractMax()
	h.ExtractMax()
	h.Insert(71)
	h.Insert(17)
	fmt.Println(h.data, h.size)
}
