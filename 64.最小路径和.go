/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// 动态规划+滚动数组

// 61/61 cases passed (4 ms)
// Your runtime beats 98.41 % of golang submissions
// Your memory usage beats 98.88 % of golang submissions (3.7 MB)

package main

import "fmt"

// @lc code=start
func minPathSum(grid [][]int) int {
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, len(grid[0])), make([]int, len(grid[0]))

	k := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 || j == 0 {
				if i == 0 && j == 0 {
					dp[k][j] = grid[i][j]
				} else if i == 0 && j > 0 {
					dp[k][j] = grid[i][j] + dp[k][j-1]
				} else {
					dp[k][j] = grid[i][j] + dp[k^1][j]
				}
			} else {
				dp[k][j] = min(dp[k^1][j], dp[k][j-1]) + grid[i][j]
			}
		}
		k ^= 1
	}
	return dp[k^1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	// fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
