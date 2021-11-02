/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func deleteNode(node *ListNode) {
	temp := node.Next

	node.Val = node.Next.Val
	node.Next = node.Next.Next
	temp.Next = nil
}

// @lc code=end
