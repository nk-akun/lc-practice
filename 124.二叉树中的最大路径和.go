/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

package main

// 94/94 cases passed (24 ms)
// Your runtime beats 7.24 % of golang submissions
// Your memory usage beats 99.04 % of golang submissions (7.1 MB)

// 递归判断
// 判断一定以node点为公共祖先的两点路径最大值能否更新ans

import "math"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func maxPathSum(root *TreeNode) int {
	ans, _ := dfs(root)
	return ans
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return math.MinInt, 0
	}
	ansL, maxL := dfs(node.Left)
	ansR, maxR := dfs(node.Right)

	return max(max(ansL, ansR), maxL+maxR+node.Val), max(0, max(maxL, maxR)+node.Val) // 小于0的直接抛弃
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
