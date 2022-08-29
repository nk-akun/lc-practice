package main

import "fmt"

func solve(s string) []int {
	result := []int{}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			result = append(result, 1)
		} else {
			if s[i] > s[i-1] {
				result = append(result, result[i-1]+1)
			} else {
				result = append(result, 1)
			}
		}
	}

	return result
}

func main() {
	var t int
	fmt.Scan(&t)

	var n int
	var s string
	ans := [][]int{}
	for t > 0 {
		t--
		fmt.Scan(&n)
		fmt.Scan(&s)
		ans = append(ans, solve(s))
	}

	for i := range ans {
		fmt.Printf("Case #%d:", i+1)
		for j := range ans[i] {
			fmt.Printf(" %d", ans[i][j])
		}
		fmt.Println()
	}
}
