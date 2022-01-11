/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// 无需头结点,nb

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var ans, p *ListNode
	if list1.Val > list2.Val {
		ans = list2
		p = list2
		list2 = list2.Next
	} else {
		ans = list1
		p = list1
		list1 = list1.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 == nil {
		p.Next = list2
	} else {
		p.Next = list1
	}
	return ans
}

// @lc code=end
