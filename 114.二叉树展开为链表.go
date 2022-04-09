/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

package main

// 根节点开始，如果当前节点有左节点，将左节点的最右节点连接到当前节点的右节点上，当然节点的右节点连接到当前节点的左节点上
// 如果没有左节点，则把当前节点切换到右节点(即链表next)

// 225/225 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 52.19 % of golang submissions (2.8 MB)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func flatten(root *TreeNode) {
	p := root
	for p != nil {
		if p.Left != nil {
			tmp := p.Left
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = p.Right
			p.Right = p.Left
			p.Left = nil
		} else {
			p = p.Right
		}
	}
}

// @lc code=end
