/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

package main

// 203/203 cases passed (4 ms)
// Your runtime beats 93.22 % of golang submissions
// Your memory usage beats 68.78 % of golang submissions (4 MB)

// pre[0]是当前子树的根，然后在in中找到pre[0]，分割左右子树
// 递归调用上述步骤即可

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	nowRoot := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	var nowRootPos int
	for i := range inorder {
		if inorder[i] == nowRoot.Val {
			nowRootPos = i
			break
		}
	}

	nowRoot.Left = buildTree(preorder[1:1+nowRootPos], inorder[:nowRootPos])
	nowRoot.Right = buildTree(preorder[1+nowRootPos:], inorder[nowRootPos+1:])
	return nowRoot
}

// @lc code=end
