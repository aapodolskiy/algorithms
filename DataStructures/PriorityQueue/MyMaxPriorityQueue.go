package main

import (
	"fmt"
	"strconv"
)

//////////////////////////////////////////////
////////// MyMaxPriorityQueue implementation

type QueueObject struct {
	s string
}

type MyMaxPriorityQueueInterface interface {
	IsEmpty() bool
	Insert(object QueueObject, priority int)
	GetObjectWithHighestPriority() QueueObject
	GetHighestPriority() int
	ExtractObjectWithHighestPriority() QueueObject
}

type MyMaxPriorityQueue struct {
	objects    *[]QueueObject
	priorities *[]int
	size       int
}

func (this *MyMaxPriorityQueue) _left(i int) int {
	return 2*(i+1) - 1
}
func (this *MyMaxPriorityQueue) _right(i int) int {
	return 2 * (i + 1)
}
func (this *MyMaxPriorityQueue) _parent(i int) int {
	return (i+1)/2 - 1
}
func (this *MyMaxPriorityQueue) _swap(i, j int) {
	(*this.objects)[i], (*this.objects)[j] = (*this.objects)[j], (*this.objects)[i]
	(*this.priorities)[i], (*this.priorities)[j] = (*this.priorities)[j], (*this.priorities)[i]
}

func (this *MyMaxPriorityQueue) _maxHeapify(i int) {
	l, r := this._left(i), this._right(i)
	largest := i
	if l < this.size && (*this.priorities)[largest] < (*this.priorities)[l] {
		largest = l
	}
	if r < this.size && (*this.priorities)[largest] < (*this.priorities)[r] {
		largest = r
	}
	if largest != i {
		this._swap(i, largest)
		this._maxHeapify(largest)
	}
}

func MyMaxPriorityQueueConstructor() MyMaxPriorityQueueInterface {
	size := 0
	objects := make([]QueueObject, size)
	priorities := make([]int, size)
	return &MyMaxPriorityQueue{&objects, &priorities, size}
}

func (this *MyMaxPriorityQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyMaxPriorityQueue) GetObjectWithHighestPriority() QueueObject {
	if this.IsEmpty() {
		var stub QueueObject
		return stub
	}
	return (*this.objects)[0]
}

func (this *MyMaxPriorityQueue) GetHighestPriority() int {
	if this.IsEmpty() {
		panic("calling GetHighestPriority on empty PQ")
	}
	return (*this.priorities)[0]
}

func (this *MyMaxPriorityQueue) ExtractObjectWithHighestPriority() QueueObject {
	if this.IsEmpty() {
		var stub QueueObject
		return stub
	}

	objectWithHighestPriority := (*this.objects)[0]

	this.size--
	this._swap(0, this.size)
	this._maxHeapify(0)

	return objectWithHighestPriority
}

func (this *MyMaxPriorityQueue) Insert(object QueueObject, priority int) {
	if len(*this.priorities) > this.size {
		(*this.priorities)[this.size] = priority
		(*this.objects)[this.size] = object
	} else {
		*this.priorities = append(*this.priorities, priority)
		*this.objects = append(*this.objects, object)
	}
	for i := this.size; i > 0 && (*this.priorities)[this._parent(i)] < (*this.priorities)[i]; i = this._parent(i) {
		this._swap(i, this._parent(i))
	}
	this.size++
}

func (this *MyMaxPriorityQueue) String() string {
	res := "priority queue: "
	for i := 0; i < this.size; i++ {
		res += fmt.Sprintf("%v", (*this.objects)[i]) +
			"[" + strconv.Itoa((*this.priorities)[i]) + "] "
	}
	return res
}

////////// MyMaxHeap implementation //////////
//////////////////////////////////////////////

func main() {
	pq := MyMaxPriorityQueueConstructor()
	pq.Insert(QueueObject{"five"}, 5)
	fmt.Println(pq.GetObjectWithHighestPriority(), pq.GetHighestPriority())
	fmt.Println(pq)

	pq.Insert(QueueObject{"seven"}, 7)
	pq.Insert(QueueObject{"seventeen"}, 17)
	pq.Insert(QueueObject{"one more seventeen"}, 17)
	fmt.Println(pq)
	fmt.Println(pq.ExtractObjectWithHighestPriority())
	pq.Insert(QueueObject{"ten"}, 10)
	fmt.Println(pq)
}
