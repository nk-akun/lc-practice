/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

package main

// 124/124 cases passed (8 ms)
// Your runtime beats 32.96 % of golang submissions
// Your memory usage beats 94.9 % of golang submissions (4.8 MB)

// 简单树形dp 时间O(N)
// 每个节点有两种情况，打劫或不打劫
// 该节点如果要打劫，子节点就不能打劫；该节点如果不打劫，子节点大不大都行
// 所以dfs返回两个值，打劫的值和不打劫的值

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func rob(root *TreeNode) int {
	ans1, ans2 := dfs(root)
	return max(ans1, ans2)
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	l1, l2 := dfs(node.Left)
	r1, r2 := dfs(node.Right)
	return l2 + r2 + node.Val, max(l1, l2) + max(r1, r2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
