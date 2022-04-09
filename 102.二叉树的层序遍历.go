/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 队列实现层次遍历即可

// 34/34 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 56.74 % of golang submissions (2.6 MB)

// @lc code=start

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	qu := []*TreeNode{root}
	for len(qu) > 0 {
		length := len(qu)
		result := []int{}
		for i := 0; i < length; i++ {
			now := qu[0]
			result = append(result, now.Val)
			if now.Left != nil {
				qu = append(qu, now.Left)
			}
			if now.Right != nil {
				qu = append(qu, now.Right)
			}
			qu = qu[1:]
		}
		ans = append(ans, result)
	}
	return ans
}

// @lc code=end
