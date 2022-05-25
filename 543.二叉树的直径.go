/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

package main

// 104/104 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 25.97 % of golang submissions (4.2 MB)

// dfs搞一搞

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		ans = max(ans, leftH+rightH)
		return max(leftH, rightH) + 1
	}

	dfs(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
