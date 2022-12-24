package main

import (
	"fmt"
)

func gcd(m, n int) int {
	if m < n {
		return gcd(n, m)
	}
	if n == 0 {
		return m
	}
	if n == 1 {
		return 1
	}
	return gcd(n, m%n)
}

func main() {
	tests := [][2]int{
		[2]int{3, 5},
		[2]int{13, 5},
		[2]int{123, 325},
		[2]int{43, 935},
		[2]int{1001, 65},
		[2]int{911, 52},
		[2]int{33, 819},
		[2]int{323212138, 597827632},
		[2]int{12 * 90 * 13, 4 * 10 * 73 * 33},
		[2]int{2 * 3 * 7 * 12192, 7 * 9 * 33120},
	}

	fmt.Println("a", "\t", "b", "\t", "gcd")
	for _, test := range tests {
		a, b := test[0], test[1]
		fmt.Println(a, "\t", b, "\t", gcd(a, b))
	}
}
