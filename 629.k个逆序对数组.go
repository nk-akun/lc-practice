/*
 * @lc app=leetcode.cn id=629 lang=golang
 *
 * [629] K个逆序对数组
 */

package main

import "fmt"

// @lc code=start

// 超时代码
// func kInversePairs(n int, k int) int {
// 	if k == 0 {
// 		return 1
// 	}

// 	mod := 1000000007
// 	dp := make([][]int, 1001)
// 	for i := range dp {
// 		dp[i] = make([]int, 1001)
// 	}

// 	for i := 1; i <= 1000; i++ {
// 		dp[i][0] = 1
// 		for j := 1; j <= 1000; j++ {
// 			for l := 0; l <= min(i-1, j); l++ {
// 				dp[i][j] = (dp[i][j] + dp[i-1][j-l]) % mod
// 			}
// 			if i == n && j == k {
// 				return dp[i][j]
// 			}
// 		}
// 	}

// 	return 0
// }

func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}

	mod := 1000000007
	dp := make([][]int, 1001)
	for i := range dp {
		dp[i] = make([]int, 1001)
	}

	for i := 1; i <= 1000; i++ {
		dp[i][0] = 1
		for j := 1; j <= 1000; j++ {
			p := j - min(i-1, j)
			var sum int
			if p > 0 {
				sum = (dp[i-1][j] - dp[i-1][p-1] + mod) % mod
			} else {
				sum = dp[i-1][j]
			}
			dp[i][j] = (dp[i][j] + sum) % mod
			if i == n && j == k {
				return dp[i][j]
			}
			dp[i][j] = (dp[i][j] + dp[i][j-1]) % mod
		}
	}

	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(kInversePairs(4, 2))
}
