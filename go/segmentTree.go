package main

import "fmt"

///////////////////////////////

func greaterOrEqualPowerOfTwo(n int) int {
	p := 1
	for p < n {
		p = p << 1
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
) {
	if left == right {
		(*this.data)[root] = (*input)[left]
		return
	}

	middle := (left + right) / 2
	leftChild, rightChild := 2*root, 2*root+1

	this._fill(input, left, middle, leftChild)
	this._fill(input, middle+1, right, rightChild)

	(*this.data)[root] = this.function(
		(*this.data)[leftChild],
		(*this.data)[rightChild],
	)
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

	leftChild, rightChild := 2*root, 2*root+1
	segmentMiddle := (segmentLeft + segmentRight) / 2

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

func (this *SegmentTree) elementUpdate(index int, newValue int) int {
	return this._recursiveElementUpdate(
		index, newValue,
		1,
		0, this.n-1,
	)
}

func (this *SegmentTree) _recursiveElementUpdate(
	index int, newValue int,
	root int,
	segmentLeft, segmentRight int,
) int {
	if index < segmentLeft || segmentRight < index {
		return (*this.data)[root]
	}
	if segmentLeft == index && index == segmentRight {
		(*this.data)[root] = newValue
		return (*this.data)[root]
	}

	segmentMiddle := (segmentLeft + segmentRight) / 2
	leftChild, rightChild := 2*root, 2*root+1

	(*this.data)[root] = this.function(
		this._recursiveElementUpdate(
			index, newValue,
			leftChild,
			segmentLeft, segmentMiddle,
		),
		this._recursiveElementUpdate(
			index, newValue,
			rightChild,
			segmentMiddle+1, segmentRight,
		),
	)
	return (*this.data)[root]
}

// only for min/max ST
func (this *SegmentTree) rangeUpdate(left, right int, newValue int) int {
	return this._recursiveRangeUpdate(
		left, right, newValue,
		1,
		0, this.n-1,
	)
}

func (this *SegmentTree) _recursiveRangeUpdate(
	queryLeft, queryRight int, newValue int,
	root int,
	segmentLeft, segmentRight int,
) int {
	if queryRight < segmentLeft || segmentRight < queryLeft {
		this._propagate(root, segmentLeft, segmentRight)
		return (*this.data)[root]
	}
	if queryLeft <= segmentLeft && segmentRight <= queryRight {
		(*this.data)[root] = newValue
		if segmentLeft < segmentRight {
			(*this.lazy)[2*root] = newValue
			(*this.lazy)[2*root+1] = newValue
		}
		return (*this.data)[root]
	}

	this._propagate(root, segmentLeft, segmentRight)

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
	maxST := createSegmentTree(&a, max, -2147483648)
	minST := createSegmentTree(&a, min, 2147483647)
	sumST := createSegmentTree(&a, sum, 0)

	maxST.elementUpdate(2, 15)
	minST.elementUpdate(2, -3)
	sumST.elementUpdate(2, 1)

	fmt.Println(maxST.data)

	maxST.rangeUpdate(0, 2, 6)
	fmt.Println(maxST.data, maxST.lazy)
	fmt.Println(maxST.rangeQuery(1, 3))
	fmt.Println(maxST.data, maxST.lazy)
	fmt.Println(maxST.rangeQuery(0, 0))
	fmt.Println(maxST.data, maxST.lazy)
}
