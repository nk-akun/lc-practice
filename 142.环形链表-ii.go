/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

package main

// 16/16 cases passed (8 ms)
// Your runtime beats 25.85 % of golang submissions
// Your memory usage beats 68.87 % of golang submissions (3.5 MB)

// 快慢指针跑一跑
// 快慢相遇时，再让second和head同时跑直到相遇为止

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	first := head
	second := head

	for {
		if first == nil || first.Next == nil {
			return nil
		}
		first = first.Next.Next
		second = second.Next
		if first == second {
			break
		}
	}

	p := head
	for p != second {
		p = p.Next
		second = second.Next
	}

	return p
}

// @lc code=end
