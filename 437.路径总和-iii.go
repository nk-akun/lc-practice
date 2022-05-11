/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */

package main

// 127/127 cases passed (4 ms)
// Your runtime beats 96.93 % of golang submissions
// Your memory usage beats 18.83 % of golang submissions (5.1 MB)

// 前缀和,map计数 时间O(N)
// 使用map记录走过的前缀和出现的次数，走到一个点时，计算以该点结尾的路径和为targetSum的数量，即map[sum-targetSum]有多少个

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func pathSum(root *TreeNode, targetSum int) int {
	valMap := map[int]int{}
	valMap[0] = 1 // 相当于0为根节点
	ans := 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		if v := valMap[sum-targetSum]; v > 0 {
			ans += v
		}
		valMap[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		valMap[sum]--
	}

	dfs(root, 0)
	return ans
}

// @lc code=end
