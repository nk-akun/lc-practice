/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// 动态规划+滚动数组
// dp[i][j] = dp[i-1][j]+dp[i][j-1]

// 62/62 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.8 MB)

package main

import "fmt"

// @lc code=start
func uniquePaths(m int, n int) int {
	ans := make([][]int, 2)
	ans[0], ans[1] = make([]int, n), make([]int, n)

	k := 0
	for i := 0; i < m; i++ {
		ans[k][0] = 1
		for j := 1; j < n; j++ {
			if i == 0 {
				ans[k][j] = 1
			} else {
				ans[k][j] = ans[k][j-1] + ans[k^1][j]
			}
		}
		k ^= 1
	}
	return ans[k^1][n-1]
}

// @lc code=end

func main() {
	fmt.Println(uniquePaths(10, 100))
}
