package main

import (
	"fmt"
)

func solve(n int) int {
	result := 0

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			result += i
			n /= i
		}
	}

	if n > 1 {
		result += n
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}
