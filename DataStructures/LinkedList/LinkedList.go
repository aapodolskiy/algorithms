package main

import (
	"fmt"
	"strconv"
)

////////////////////////////////
// implementation of MyLinkedList

type MyLinkedListInterface interface {
	Get(index int) int
	AddAtHead(val int)
	AddAtTail(val int)
	AddAtIndex(index int, val int)
	DeleteAtIndex(index int)
}

type MyListNode struct {
	Val  int
	Next *MyListNode
}

type MyLinkedList struct {
	head *MyListNode
}

func MyLinkedListConstructor() MyLinkedListInterface {
	var head *MyListNode
	return &MyLinkedList{head}
}

func (this *MyLinkedList) Get(index int) int {
	for node := this.head; index >= 0; index, node = index-1, node.Next {
		if node == nil {
			break
		}
		if index == 0 {
			return node.Val
		}
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &MyListNode{val, this.head}
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := this.head
	if node == nil {
		this.AddAtHead(val)
		return
	}
	for ; node.Next != nil; node = node.Next {
	}
	node.Next = &MyListNode{val, nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
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
		res += strconv.Itoa(node.Val) + " -> "
	}
	return res + "nil"
}

// end of MyLinkedList implementation
////////////////////////////////

func main() {
	obj := MyLinkedListConstructor()
	fmt.Println(obj.Get(0))
	obj.AddAtHead(3)
	obj.AddAtTail(4)
	obj.AddAtIndex(1, 5)
	obj.DeleteAtIndex(0)
	fmt.Println(obj)
}
