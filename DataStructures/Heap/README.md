# Heap

Heap is a specialized tree-based data structure which is essentially an almost complete tree that satisfies the heap property: in a max heap, for any given node C, if P is a parent node of C, then the key (the value) of P is greater than or equal to the key of C. In a min heap, the key of P is less than or equal to the key of C. The node at the "top" of the heap (with no parents) is called the root node.


This file (it will be module in future) should export two functions:

```go
func MyMaxHeapConstructor(input ...int) MyHeapInterface

func MyMinHeapConstructor(input ...int) MyHeapInterface
```


Both functions return an object, which implements `MyHeapInterface`:

```go
type MyHeapInterface interface {
	IsEmpty() bool
	GetSize() int
	Insert(int)
	GetTop() int
	ExtractTop() int
}
```

## Functions


### `func MyMaxHeapConstructor(input ...int) MyHeapInterface`

#### Description

returns heap object constructed from array of integers (possibly empty)

#### example

```go
myMaxHeap := MyMaxHeapConstructor()

myMinHeap := MyMinHeapConstructor([]int{ 5, 3, 8 }...)
```

#### complexity

$O(N)$ in worst case


### `func MyMinHeapConstructor(input ...int) MyHeapInterface`

as above, but min-heap property is hold


## Methods



### `IsEmpty() bool`

#### Description

returns `true` if Heap is empty, `false` otherwise

#### example

```go
myMaxHeap := MyMaxHeapConstructor()
myMaxHeap.IsEmpty() // true

myMinHeap := MyMinHeapConstructor([]int{ 5, 3, 8 }...)
myMinHeap.IsEmpty() // false
```

#### complexity

$O(1)$ in worst case



### `GetSize() int`

#### Description

returns actual number of elements in heap. 

#### example

```go
myMaxHeap := MyMaxHeapConstructor()
myMaxHeap.GetSize() // 0
myMaxHeap.Insert(2)
myMaxHeap.GetSize() // 1

myMinHeap := MyMinHeapConstructor([]int{ 5, 3, 8 }...)
myMinHeap.GetSize() // 3
myMinHeap.ExtractTop()
myMinHeap.GetSize() // 2
```

#### complexity

$O(1)$ in worst case



### `Insert(int)`

#### Description

inserts an element into heap so that heap property is still satisfied

#### example

```go
myMaxHeap := MyMaxHeapConstructor()
for _, e := range []int{ 2, 7, 3, 4 } {
    myMaxHeap.Insert(e*e)
}
```

#### complexity

$O(\log N)$ in worst case 


---

TODO

- one more method
- check how it looks on github

<details>
  <summary>Click me</summary>
  
  ### Heading
  1. Foo
  2. Bar
     * Baz
     * Qux

  ### Some Code
  ```js
  function logSomething(something) {
    console.log('Something', something);
  }
  ```
</details>

<img src="https://latex.codecogs.com/svg.latex?\Large&space;O(n)" title="Complexity" />