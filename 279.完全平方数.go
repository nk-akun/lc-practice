/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

package main

// 588/588 cases passed (24 ms)
// Your runtime beats 64.98 % of golang submissions
// Your memory usage beats 89.8 % of golang submissions (6 MB)

// 动态规划
// dp跑一跑

import (
	"fmt"
	"math"
)

// @lc code=start
func numSquares(n int) int {
	quares := []int{}
	for i := 1; i*i <= n; i++ {
		quares = append(quares, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < len(quares) && quares[j] <= i; j++ {
			if dp[i-quares[j]]+1 < dp[i] {
				dp[i] = dp[i-quares[j]] + 1
			}
		}
	}

	return dp[n]
}

// @lc code=end

func main() {
	fmt.Println(numSquares(1028))
}
