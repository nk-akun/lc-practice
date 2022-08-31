package main

import "fmt"

const mod = 1000000007

func solve(n int) int {
	dp := make([][]int, n+1)

	dp[1] = make([]int, 3)
	dp[1][0] = 8
	dp[1][1] = 1
	dp[1][2] = 0
	for i := 2; i <= n; i++ {
		dp[i] = make([]int, 3)
		dp[i][0] = ((dp[i-1][0] + dp[i-1][2]) * 8) % mod
		dp[i][1] = (dp[i-1][0] * 1) % mod
		dp[i][2] = (dp[i-1][1] * 8) % mod
	}

	ans := 0
	for i := 0; i < 3; i++ {
		ans = (ans + dp[n][i]) % mod
	}
	return ans
}

// 对拍
var sum int

func judge(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != 1 {
			continue
		}
		if a[i-1] == 1 || i > 1 && a[i-2] == 1 {
			return false
		}
	}
	return true
}

func dfs(a []int, pos int) {
	if pos == len(a) {
		if judge(a) {
			sum++
		}
		return
	}
	for i := 1; i <= 9; i++ {
		a[pos] = i
		dfs(a, pos+1)
	}
}

func solve2(n int) int {
	a := make([]int, n)
	dfs(a, 0)
	return sum
}

// 对拍

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(solve(n))
	fmt.Println(solve2(n))
}
