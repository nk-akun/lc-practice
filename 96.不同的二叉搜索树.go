/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

package main

// 递归来解即可

// 19/19 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 8.64 % of golang submissions (1.9 MB)

import "fmt"

// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
	var dfs func(int) int
	dfs = func(x int) int {
		if x <= 1 {
			return 1
		}
		if dp[x] > 0 {
			return dp[x]
		}
		result := 0
		for i := 1; i <= x; i++ {
			result += dfs(i-1) * dfs(x-i)
		}
		dp[x] = result
		return result
	}

	return dfs(n)
}

// @lc code=end

func main() {
	fmt.Println(numTrees(4))
}
