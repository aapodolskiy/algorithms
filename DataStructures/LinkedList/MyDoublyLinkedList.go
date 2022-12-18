package main

import (
	"fmt"
)

////////////////////////////////
// implementation of MyDoublyLinkedList

type MyListValue struct {
	i int
}

type MyDoublyLinkedListInterface interface {
	Size() int
	GetHead() *MyListNode
	GetTail() *MyListNode
	Get(p *MyListNode) MyListValue
	UpdateVal(p *MyListNode, val MyListValue)
	AddAtHead(val MyListValue) *MyListNode
	AddAtTail(val MyListValue) *MyListNode
	Delete(p *MyListNode)
}

type MyListNode struct {
	Val  MyListValue
	Prev *MyListNode
	Next *MyListNode
}

type MyDoublyLinkedList struct {
	head *MyListNode
	tail *MyListNode
	size int
}

func MyDoublyLinkedListConstructor() MyDoublyLinkedListInterface {
	var stub MyListValue
	head := &MyListNode{stub, nil, nil}
	tail := &MyListNode{stub, head, nil}
	head.Next = tail
	return &MyDoublyLinkedList{head, tail, 0}
}

func (this *MyDoublyLinkedList) Size() int {
	return this.size
}

func (this *MyDoublyLinkedList) GetHead() *MyListNode {
	return this.head
}

func (this *MyDoublyLinkedList) GetTail() *MyListNode {
	return this.tail
}

func (this *MyDoublyLinkedList) Get(p *MyListNode) MyListValue {
	for p == nil {
		var stub MyListValue
		return stub
	}
	return p.Val
}

func (this *MyDoublyLinkedList) UpdateVal(p *MyListNode, val MyListValue) {
	for p == nil || p == this.head || p == this.tail {
		return
	}
	p.Val = val
}

func (this *MyDoublyLinkedList) AddAtHead(val MyListValue) *MyListNode {
	newNode := &MyListNode{val, this.head, this.head.Next}
	this.head.Next.Prev = newNode
	this.head.Next = newNode
	this.size++
	return newNode
}

func (this *MyDoublyLinkedList) AddAtTail(val MyListValue) *MyListNode {
	newNode := &MyListNode{val, this.tail.Prev, this.tail}
	this.tail.Prev.Next = newNode
	this.tail.Prev = newNode
	this.size++
	return newNode
}

func (this *MyDoublyLinkedList) Delete(p *MyListNode) {
	if p == this.head || p == this.tail {
		return
	}
	prev, next := p.Prev, p.Next
	prev.Next = next
	next.Prev = prev
	this.size--
}

func (this *MyDoublyLinkedList) String() string {
	res := "doubly linked list: head <-> "
	for node := this.head.Next; node != this.tail; node = node.Next {
		res += fmt.Sprintf("%v", node.Val) + " <-> "
	}
	res += "tail"
	return res
}

// end of MyDoublyLinkedList implementation
////////////////////////////////

func main() {
	obj := MyDoublyLinkedListConstructor()
	obj.AddAtHead(MyListValue{3})
	four := obj.AddAtHead(MyListValue{4})
	obj.AddAtHead(MyListValue{5})
	eleven := obj.AddAtTail(MyListValue{11})
	obj.AddAtTail(MyListValue{12})
	fmt.Println(obj)

	obj.Delete(eleven)
	fmt.Println(obj)

	fmt.Println(obj.Get(four))

	first := obj.AddAtHead(eleven.Val)
	fmt.Println(obj)

	obj.UpdateVal(first, MyListValue{1})
	fmt.Println(obj)

	fmt.Println(obj.Size())

	fmt.Println(obj.GetHead().Next)
	fmt.Println(obj.GetTail().Prev)
}
