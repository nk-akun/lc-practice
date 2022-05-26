/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */

package main

// 215/215 cases passed (8 ms)
// Your runtime beats 94.91 % of golang submissions
// Your memory usage beats 99.23 % of golang submissions (6.6 MB)

// 反向中序遍历，先后再中再左
// 这样就能直接 把目前遍历过所有的节点的值的和 累加在 当前节点上 即可

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}

	dfs(root)
	return root
}

// @lc code=end
