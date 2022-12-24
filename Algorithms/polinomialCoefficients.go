package main

import (
	"fmt"
)

func polinomialCoefficient(a []int, primeModulus int) int {
	res := 1
	sum := 0
	for _, v := range a {
		sum += v
		res *= binomialCoefficient(sum, v, primeModulus)
		res %= primeModulus
	}
	return res
}

func binomialCoefficient(
	n, k int,
	primeModulus int,
) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k == 1 {
		return n % primeModulus
	}

	inverse := make([]int, k+1)
	inverse[0] = 1
	inverse[1] = 1
	for i := 2; i <= k; i++ {
		inverse[i] = -(primeModulus / i) * inverse[primeModulus%i]
		inverse[i] %= primeModulus
		inverse[i] += primeModulus
		inverse[i] %= primeModulus
	}

	res := 1
	for i := 2; i <= k; i++ {
		res *= inverse[i] % primeModulus
		res %= primeModulus
	}
	for i := n; i >= n-k+1; i-- {
		res *= i % primeModulus
		res %= primeModulus
	}
	return res
}

func main() {
	tests := [][]int{
		[]int{1, 1, 1},
		[]int{1, 2},
		[]int{2, 1},
		[]int{2, 3, 4},
		[]int{3, 5, 1, 2, 3},
		[]int{911, 52, 764, 21, 9, 4},
	}

	M := 1000000007

	for _, test := range tests {
		fmt.Println(test, "\t", polinomialCoefficient(test, M))
	}
}
