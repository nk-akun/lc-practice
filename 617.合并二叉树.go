/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
 */

package main

// 182/182 cases passed (12 ms)
// Your runtime beats 91.18 % of golang submissions
// Your memory usage beats 51.71 % of golang submissions (6.7 MB)

// 通过dfs把tree1往tree2里合
// 如果tree1当前节点和tree2当前节点都有左子树，那dfs左子树
// 如果tree1有左子树，tree2没有，那么tree2的当前节点的左节点指向tree1的左节点
// 其他情况不用处理（tree1没有左子树，tree2有左子树 或者 两者均没有）
// 右子树同理
// return root2即可

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	var dfs func(node1 *TreeNode, node2 *TreeNode)
	dfs = func(node1 *TreeNode, node2 *TreeNode) {
		node2.Val += node1.Val
		if node2.Left == nil && node1.Left != nil {
			node2.Left = node1.Left
		} else if node1.Left != nil && node2.Left != nil {
			dfs(node1.Left, node2.Left)
		}

		if node2.Right == nil && node1.Right != nil {
			node2.Right = node1.Right
		} else if node1.Right != nil && node2.Right != nil {
			dfs(node1.Right, node2.Right)
		}
	}

	dfs(root1, root2)
	return root2
}

// @lc code=end
