package main

import (
	"fmt"
)

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
	tests := [][2]int{
		[2]int{3, 5},
		[2]int{13, 5},
		[2]int{123, 25},
		[2]int{43, 11},
		[2]int{1001, 65},
		[2]int{911, 52},
		[2]int{3033, 819},
		[2]int{323212138, 527632},
	}

	M := 1000000007

	fmt.Println("n", "\t", "k", "\t", "binomialCoefficient(n, k) mod M")
	for _, test := range tests {
		n, k := test[0], test[1]
		fmt.Println(n, "\t", k, "\t", binomialCoefficient(n, k, M))
	}
}
