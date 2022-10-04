package main

import "fmt"

///////////////////////////////

func greaterOrEqualPowerOfTwo(n int) int {
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

type SegmentTree struct {
	n            int    // number of segment borders
	data         *[]int // ST values
	lazy         *[]int // for lazy propagation
	function     func(int, int) int
	defaultValue int
}

func createSegmentTree(
	input *[]int,
	f func(int, int) int,
	defaultValue int,
) SegmentTree {
	n := len(*input)
	data := make([]int, 2*greaterOrEqualPowerOfTwo(n))
	lazy := make([]int, 2*greaterOrEqualPowerOfTwo(n))
	ST := SegmentTree{n, &data, &lazy, f, defaultValue}
	ST._fill(
		input, 0, n-1,
		1,
	)
	return ST
}

func (this *SegmentTree) _fill(
	input *[]int, left, right int,
	root int,
) int {
	if left == right {
		(*this.data)[root] = (*input)[left]
		return (*this.data)[root]
	}

	middle := (left + right) / 2
	leftChild, rightChild := 2*root, 2*root+1

	(*this.data)[root] = this.function(
		this._fill(input, left, middle, leftChild),
		this._fill(input, middle+1, right, rightChild),
	)
	return (*this.data)[root]
}

func (this *SegmentTree) rangeQuery(left, right int) int {
	return this._recursiveRangeQuery(
		left, right,
		1,
		0, this.n-1,
	)
}

func (this *SegmentTree) _propagate(
	root int,
	segmentLeft, segmentRight int,
) {
	if (*this.lazy)[root] == 0 {
		return
	}
	newValue := (*this.lazy)[root]
	(*this.lazy)[root] = 0
	(*this.data)[root] = newValue
	if segmentLeft < segmentRight {
		leftChild, rightChild := 2*root, 2*root+1
		(*this.lazy)[leftChild] = newValue
		(*this.lazy)[rightChild] = newValue
	}
}

func (this *SegmentTree) _recursiveRangeQuery(
	queryLeft, queryRight int,
	root int,
	segmentLeft, segmentRight int,
) int {
	if queryRight < segmentLeft || segmentRight < queryLeft {
		return this.defaultValue
	}

	this._propagate(root, segmentLeft, segmentRight)

	if queryLeft <= segmentLeft && segmentRight <= queryRight {
		return (*this.data)[root]
	}

	segmentMiddle := (segmentLeft + segmentRight) / 2
	leftChild, rightChild := 2*root, 2*root+1

	return this.function(
		this._recursiveRangeQuery(
			queryLeft, queryRight,
			leftChild,
			segmentLeft, segmentMiddle,
		),
		this._recursiveRangeQuery(
			queryLeft, queryRight,
			rightChild,
			segmentMiddle+1, segmentRight,
		),
	)
}

// not for sum ST
func (this *SegmentTree) rangeUpdate(left, right int, newValue int) int {
	return this._recursiveRangeUpdate(
		left, right, newValue,
		1,
		0, this.n-1,
	)
}

func (this *SegmentTree) _delayPropagation(
	newValue int,
	root int,
	segmentLeft, segmentRight int,
) {
	(*this.data)[root] = newValue
	if segmentLeft < segmentRight {
		(*this.lazy)[2*root] = newValue
		(*this.lazy)[2*root+1] = newValue
	}
}

func (this *SegmentTree) _recursiveRangeUpdate(
	queryLeft, queryRight int, newValue int,
	root int,
	segmentLeft, segmentRight int,
) int {
	this._propagate(root, segmentLeft, segmentRight)

	if queryRight < segmentLeft || segmentRight < queryLeft {
		return (*this.data)[root]
	}

	if queryLeft <= segmentLeft && segmentRight <= queryRight {
		this._delayPropagation(newValue, root, segmentLeft, segmentRight)
		return (*this.data)[root]
	}

	segmentMiddle := (segmentLeft + segmentRight) / 2
	leftChild, rightChild := 2*root, 2*root+1

	(*this.data)[root] = this.function(
		this._recursiveRangeUpdate(
			queryLeft, queryRight, newValue,
			leftChild,
			segmentLeft, segmentMiddle,
		),
		this._recursiveRangeUpdate(
			queryLeft, queryRight, newValue,
			rightChild,
			segmentMiddle+1, segmentRight,
		),
	)
	return (*this.data)[root]
}

///////////////////////////////

func main() {
	a := []int{1, 2, 9, 7, 8, 9}
	max := func(x int, y int) int {
		if x < y {
			return y
		}
		return x
	}
	min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}
	sum := func(x int, y int) int {
		return x + y
	}

	fmt.Println("maxST")
	maxST := createSegmentTree(&a, max, -2147483648)
	maxST.rangeUpdate(1, 3, 1)
	fmt.Println(maxST.data, maxST.lazy)
	fmt.Println(maxST.rangeQuery(0, 4))
	fmt.Println(maxST.data, maxST.lazy)
	fmt.Println(maxST.rangeQuery(3, 3))
	fmt.Println(maxST.data, maxST.lazy)
	fmt.Println("\n")

	fmt.Println("minST")
	minST := createSegmentTree(&a, min, 2147483647)
	minST.rangeUpdate(1, 3, 1)
	fmt.Println(minST.data, minST.lazy)
	fmt.Println(minST.rangeQuery(1, 2))
	fmt.Println(minST.rangeQuery(3, 4))
	fmt.Println(minST.data, minST.lazy)
	fmt.Println(minST.rangeQuery(0, 0))
	fmt.Println("\n")

	fmt.Println("sumST")
	sumST := createSegmentTree(&a, sum, 0)
	fmt.Println(sumST.rangeQuery(1, 3))
	fmt.Println(sumST.rangeQuery(0, 0))
	fmt.Println(sumST.rangeQuery(0, 1))
	fmt.Println(sumST.rangeQuery(0, 2))
	fmt.Println(sumST.rangeQuery(3, 3))
	fmt.Println("\n")
}
