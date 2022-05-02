/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

package main

// 21/21 cases passed (8 ms)
// Your runtime beats 38.77 % of golang submissions
// Your memory usage beats 99.98 % of golang submissions (4.2 MB)

// 快慢指针走一走

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	first := head.Next
	second := head

	for first != second {
		if first == nil || first.Next == nil {
			return false
		}
		first = first.Next.Next
		second = second.Next
	}

	return true
}

// @lc code=end
