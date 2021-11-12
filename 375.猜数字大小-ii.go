/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */

package main

import "fmt"

// @lc code=start
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = 0x3f3f3f3f
			for k := i; k <= j; k++ {
				now := max(dp[i][k-1], dp[k+1][j]) + k
				if now < dp[i][j] {
					dp[i][j] = now
				}
			}
		}
	}
	return dp[1][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(getMoneyAmount(10))
}
