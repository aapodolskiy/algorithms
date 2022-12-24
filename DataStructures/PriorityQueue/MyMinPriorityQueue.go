package main

import (
	"fmt"
	"strconv"
)

//////////////////////////////////////////////
////////// MyMinPriorityQueue implementation

type QueueObject struct {
	s string
}

type MyMinPriorityQueueInterface interface {
	IsEmpty() bool
	GetSize() int
	Insert(object QueueObject, priority int)
	GetObjectWithLowestPriority() QueueObject
	GetLowestPriority() int
	ExtractObjectWithLowestPriority() QueueObject
}

type MyMinPriorityQueue struct {
	objects    *[]QueueObject
	priorities *[]int
	size       int
}

func MyMinPriorityQueueConstructor() MyMinPriorityQueueInterface {
	size := 0
	objects := make([]QueueObject, size)
	priorities := make([]int, size)
	return &MyMinPriorityQueue{&objects, &priorities, size}
}

func (this *MyMinPriorityQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyMinPriorityQueue) GetSize() int {
	return this.size
}

func (this *MyMinPriorityQueue) GetObjectWithLowestPriority() QueueObject {
	if this.IsEmpty() {
		var stub QueueObject
		return stub
	}
	return (*this.objects)[0]
}

func (this *MyMinPriorityQueue) GetLowestPriority() int {
	if this.IsEmpty() {
		panic("calling GetLowestPriority on empty PQ")
	}
	return (*this.priorities)[0]
}

func (this *MyMinPriorityQueue) ExtractObjectWithLowestPriority() QueueObject {
	if this.IsEmpty() {
		var stub QueueObject
		return stub
	}

	objectWithLowestPriority := (*this.objects)[0]

	this.size--
	this._swap(0, this.size)
	this._minHeapify(0)

	return objectWithLowestPriority
}

func (this *MyMinPriorityQueue) Insert(object QueueObject, priority int) {
	if len(*this.priorities) > this.size {
		(*this.priorities)[this.size] = priority
		(*this.objects)[this.size] = object
	} else {
		*this.priorities = append(*this.priorities, priority)
		*this.objects = append(*this.objects, object)
	}
	for i := this.size; i > 0 && (*this.priorities)[i] < (*this.priorities)[this._parent(i)]; i = this._parent(i) {
		this._swap(i, this._parent(i))
	}
	this.size++
}

func (this *MyMinPriorityQueue) _left(i int) int {
	return 2*(i+1) - 1
}
func (this *MyMinPriorityQueue) _right(i int) int {
	return 2 * (i + 1)
}
func (this *MyMinPriorityQueue) _parent(i int) int {
	return (i+1)/2 - 1
}
func (this *MyMinPriorityQueue) _swap(i, j int) {
	(*this.objects)[i], (*this.objects)[j] = (*this.objects)[j], (*this.objects)[i]
	(*this.priorities)[i], (*this.priorities)[j] = (*this.priorities)[j], (*this.priorities)[i]
}
func (this *MyMinPriorityQueue) _minHeapify(i int) {
	l, r := this._left(i), this._right(i)
	smallest := i
	if l < this.size && (*this.priorities)[l] < (*this.priorities)[smallest] {
		smallest = l
	}
	if r < this.size && (*this.priorities)[r] < (*this.priorities)[smallest] {
		smallest = r
	}
	if smallest != i {
		this._swap(i, smallest)
		this._minHeapify(smallest)
	}
}

func (this *MyMinPriorityQueue) String() string {
	res := "min priority queue: "
	for i := 0; i < this.size; i++ {
		res += fmt.Sprintf("%v", (*this.objects)[i]) +
			"[" + strconv.Itoa((*this.priorities)[i]) + "] "
	}
	return res
}

////////// MyMinPriorityQueue implementation
//////////////////////////////////////////////

func main() {
	pq := MyMinPriorityQueueConstructor()
	pq.Insert(QueueObject{"five"}, 5)
	fmt.Println(pq.GetObjectWithLowestPriority(), pq.GetLowestPriority())
	fmt.Println(pq)

	pq.Insert(QueueObject{"seven"}, 7)
	pq.Insert(QueueObject{"seventeen"}, 17)
	pq.Insert(QueueObject{"one more seventeen"}, 17)
	fmt.Println(pq.GetSize(), pq)
	fmt.Println(pq.ExtractObjectWithLowestPriority(), pq.GetSize())
	pq.Insert(QueueObject{"ten"}, 10)
	fmt.Println(pq)
}
