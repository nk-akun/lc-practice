/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

package main

// 31/31 cases passed (8 ms)
// Your runtime beats 90.38 % of golang submissions
// Your memory usage beats 63.28 % of golang submissions (7.4 MB)

// 递归
// 优化代码有点晦涩，左右子树分别找到一个时，当前节点就是公共祖先
// 否则如果当前节点是p或q，均返回当前节点

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

// 优化代码，有点晦涩
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	_, _, ans := find(root, p, q)
// 	return ans
// }

// func find(t, p, q *TreeNode) (bool, bool, *TreeNode) {
// 	if t == nil {
// 		return false, false, nil
// 	}

// 	lexistp, lexistq, resultl := find(t.Left, p, q)
// 	rexistp, rexistq, resultr := find(t.Right, p, q)

// 	if resultl != nil {
// 		return true, true, resultl
// 	}
// 	if resultr != nil {
// 		return true, true, resultr
// 	}
// 	if lexistp && rexistq || lexistq && rexistp {
// 		return true, true, t
// 	}

// 	if t == p {
// 		if lexistq || rexistq {
// 			return true, true, t
// 		}
// 		return true, false, nil
// 	}

// 	if t == q {
// 		if lexistp || rexistp {
// 			return true, true, t
// 		}
// 		return false, true, nil
// 	}

// 	return lexistp || rexistp, lexistq || rexistq, nil
// }

// @lc code=end
