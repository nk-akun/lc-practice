/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

package main

// 188/188 cases passed (4 ms)
// Your runtime beats 99.63 % of golang submissions
// Your memory usage beats 95.29 % of golang submissions (6.2 MB)

// 完全背包,O(coins.length*amount)

import (
	"fmt"
)

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = 0x3f3f3f3f
	}
	for i := range coins {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == 0x3f3f3f3f {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(coinChange([]int{2}, 0))
}
