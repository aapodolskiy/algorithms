package main

import (
	"fmt"
)

////////////////////////////////
// implementation of MyStack

type MyStackElement string

type MyStack struct {
	a []MyStackElement
}

func NewMyStack() MyStack {
	return MyStack{make([]MyStackElement, 0)}
}

func (this *MyStack) Push(x MyStackElement) {
	this.a = append(this.a, x)
}

func (this *MyStack) Pop() MyStackElement {
	l := len(this.a)
	t := this.a[l-1]
	this.a = this.a[:l-1]
	return t
}

func (this *MyStack) Peek() MyStackElement {
	l := len(this.a)
	return this.a[l-1]
}

func (this *MyStack) Size() int {
	return len(this.a)
}

func (this *MyStack) Empty() bool {
	return len(this.a) == 0
}

// end of MyStack implementation
////////////////////////////////

func main() {
	s := NewMyStack()
	s.Push("1")
	s.Push("two")
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Size())
	fmt.Println(s.Peek())
}
