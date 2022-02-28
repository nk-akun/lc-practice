package main

// 小于等于N的所有数中有多少个1
// 经典leetcode hard

import "fmt"

func Solution(N int) int {
	dig, res := 1, 0
	high := N / 10
	current := N % 10
	low := 0
	for high != 0 || current != 0 {
		if current == 0 {
			res += high * dig
		} else if current == 1 {
			res += high*dig + low + 1
		} else {
			res += (high + 1) * dig
		}
		low += current * dig
		current = high % 10
		high /= 10
		dig *= 10
	}
	return res
}

func main() {
	fmt.Println(Solution(123451213))

	a := -1000000000
	fmt.Println(a * a)
}
