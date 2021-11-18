/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

var ans int

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = 0
	leftVal := dfs(root.Left)
	rightVal := dfs(root.Right)
	val := leftVal - rightVal
	if val < 0 {
		val = -val
	}
	ans += val
	return ans
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftVal := dfs(node.Left)
	rightVal := dfs(node.Right)
	val := leftVal - rightVal
	if val < 0 {
		val = -val
	}
	ans += val

	return leftVal + rightVal + node.Val
}

// @lc code=end

// func main() {

// }
