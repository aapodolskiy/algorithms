package main

import (
	"fmt"
	"strconv"
)

////////////////////////////////
// MyStack implementation

type MyStackInterface interface {
	Empty() bool
	Push(int)
	Pop() int
	Peek() int
}

type MyStack struct {
	a []int
}

func NewMyStack() MyStackInterface {
	return &MyStack{make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.a = append(this.a, x)
}

func (this *MyStack) Pop() int {
	l := len(this.a)
	t := this.a[l-1]
	this.a = this.a[:l-1]
	return t
}

func (this *MyStack) Peek() int {
	l := len(this.a)
	return this.a[l-1]
}

func (this *MyStack) Empty() bool {
	return len(this.a) == 0
}

func (this *MyStack) String() string {
	res := "stack:"
	for _, e := range this.a {
		res += " " + strconv.Itoa(e)
	}
	return res
}

// end of MyStack implementation
////////////////////////////////

func main() {
	s := NewMyStack()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
	fmt.Println(s.Peek())
	s.Push(7)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
}
