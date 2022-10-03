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
	treeData     *[]int // ST values
	treeFunction func(int, int) int
	defaultValue int
	// lazy *[]int // for lazy propagation
}

func createSegmentTree(
	input *[]int,
	f func(int, int) int,
	defaultValue int,
) SegmentTree {
	n := len(*input)
	treeData := make([]int, 2*greaterOrEqualPowerOfTwo(n))
	ST := SegmentTree{n, &treeData, f, defaultValue}
	ST._fill(
		input, 0, n-1,
		1,
	)
	return ST
}

func (this *SegmentTree) _fill(
	input *[]int, left, right int,
	rootIndex int,
) {
	if left == right {
		(*this.treeData)[rootIndex] = (*input)[left]
		return
	}

	middle := (left + right) / 2
	leftChild, rightChild := 2*rootIndex, 2*rootIndex+1

	this._fill(input, left, middle, leftChild)
	this._fill(input, middle+1, right, rightChild)

	(*this.treeData)[rootIndex] = this.treeFunction(
		(*this.treeData)[leftChild],
		(*this.treeData)[rightChild],
	)
}

func (this *SegmentTree) rangeQuery(left, right int) int {
	return this._recursiveRangeQuery(
		left, right,
		1,
		0, this.n-1,
	)
}

func (this *SegmentTree) _recursiveRangeQuery(
	queryLeft, queryRight int,
	root int,
	segmentLeft, segmentRight int,
) int {
	if queryRight < segmentLeft || segmentRight < queryLeft {
		return this.defaultValue
	}
	if queryLeft <= segmentLeft && segmentRight <= queryRight {
		return (*this.treeData)[root]
	}

	segmentMiddle := (segmentLeft + segmentRight) / 2
	leftChild, rightChild := 2*root, 2*root+1

	return this.treeFunction(
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

func (this *SegmentTree) rangeUpdate(
	left, right int,
	newValue int,
) int {
	return this._recursiveRangeQuery(
		left, right,
		1,
		0, this.n-1,
	)
}

///////////////////////////////

func main() {
	a := []int{1, 2, 9, 7, 3, 4}
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
	fmt.Println(maxST.treeData)
	fmt.Println(minST.treeData)
	fmt.Println(sumST.treeData)

	fmt.Println(minST.rangeQuery(1, 3))
	fmt.Println(maxST.rangeQuery(1, 3))
	fmt.Println(sumST.rangeQuery(1, 3))
	fmt.Println(minST.rangeQuery(2, 5))
	fmt.Println(maxST.rangeQuery(2, 5))
	fmt.Println(sumST.rangeQuery(2, 5))
}
