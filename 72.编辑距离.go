/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// 动态规划
// dp[i][j]表示word1的前i个字符拼成word2的前j个字符需要的操作数
// 转移方程，当word1[i]和word2[j]相同时，dp[i][j] = dp[i-1][j-1]
// 当不相同时，dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
// dp[i-1][j]表示转移到dp[i][j]需要增加一个字符,dp[i][j-1]表示删除一个字符,dp[i-1][j-1]表示替换一个字符

// 1146/1146 cases passed (4 ms)
// Your runtime beats 82.74 % of golang submissions
// Your memory usage beats 93.47 % of golang submissions (5.4 MB)

package main

import "fmt"

// @lc code=start
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(minDistance("abc", "acb"))
}
