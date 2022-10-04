package main

import "fmt"

///////////////////////////////

func greaterOrEqualPowerOfTwo(n int) (p int) {
	for p = 1; p < n; p <<= 1 {
	}
	return
}

type SegmentTree struct {
	n            int    // number of segment borders
	data         *[]int // ST values
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
	ST := SegmentTree{n, &data, f, defaultValue}
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

func (this *SegmentTree) _recursiveRangeQuery(
	queryLeft, queryRight int,
	root int,
	segmentLeft, segmentRight int,
) int {
	if queryRight < segmentLeft || segmentRight < queryLeft {
		return this.defaultValue
	}

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
	maxST.elementUpdate(2, 15)
	maxST.elementUpdate(1, 11)
	fmt.Println(maxST.data)
	fmt.Println(maxST.rangeQuery(1, 3))
	fmt.Println(maxST.rangeQuery(0, 0))
	fmt.Println("\n")

	fmt.Println("minST")
	minST := createSegmentTree(&a, min, 2147483647)
	minST.elementUpdate(2, 15)
	fmt.Println(minST.data)
	fmt.Println(minST.rangeQuery(1, 3))
	fmt.Println(minST.rangeQuery(0, 0))
	fmt.Println("\n")

	fmt.Println("sumST")
	sumST := createSegmentTree(&a, sum, 0)
	sumST.elementUpdate(2, 15)
	fmt.Println(sumST.data)
	fmt.Println(sumST.rangeQuery(1, 3))
	fmt.Println(sumST.rangeQuery(0, 0))
	fmt.Println(sumST.rangeQuery(0, 1))
	fmt.Println(sumST.rangeQuery(0, 2))
	fmt.Println(sumST.rangeQuery(3, 3))
	fmt.Println("\n")
}
