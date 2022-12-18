package main

import (
	"fmt"
)

////////////////////////////////
// implementation of MyLinkedList

type MyListValue struct {
	i int
}

type MyLinkedListInterface interface {
	Get(index int) MyListValue
	AddAtHead(val MyListValue)
	AddAtTail(val MyListValue)
	AddAtIndex(index int, val MyListValue)
	DeleteAtIndex(index int)
}

type MyListNode struct {
	Val  MyListValue
	Next *MyListNode
}

type MyLinkedList struct {
	head *MyListNode
}

func MyLinkedListConstructor() MyLinkedListInterface {
	var head *MyListNode
	return &MyLinkedList{head}
}

func (this *MyLinkedList) Get(index int) MyListValue {
	for node := this.head; index >= 0; index, node = index-1, node.Next {
		if node == nil {
			break
		}
		if index == 0 {
			return node.Val
		}
	}
	var stub MyListValue
	return stub
}

func (this *MyLinkedList) AddAtHead(val MyListValue) {
	this.head = &MyListNode{val, this.head}
}

func (this *MyLinkedList) AddAtTail(val MyListValue) {
	node := this.head
	if node == nil {
		this.AddAtHead(val)
		return
	}
	for ; node.Next != nil; node = node.Next {
	}
	node.Next = &MyListNode{val, nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val MyListValue) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	for prev := this.head; index >= 1; index, prev = index-1, prev.Next {
		if prev == nil {
			return
		}
		if index == 1 {
			prev.Next = &MyListNode{val, prev.Next}
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.head != nil {
			this.head = this.head.Next
		}
		return
	}
	for prev := this.head; index >= 1; index, prev = index-1, prev.Next {
		if prev == nil {
			return
		}
		if index == 1 && prev.Next != nil {
			prev.Next = prev.Next.Next
		}
	}
}

func (this *MyLinkedList) String() string {
	res := "linked list: "
	for node := this.head; node != nil; node = node.Next {
		res += fmt.Sprintf("%v", node.Val) + " -> "
	}
	return res + "nil"
}

// end of MyLinkedList implementation
////////////////////////////////

func main() {
	obj := MyLinkedListConstructor()
	fmt.Println(obj.Get(0))
	obj.AddAtHead(MyListValue{3})
	obj.AddAtTail(MyListValue{4})
	obj.AddAtIndex(1, MyListValue{5})
	obj.DeleteAtIndex(0)
	fmt.Println(obj)
}
