/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

package main

// 39/39 cases passed (28 ms)
// Your runtime beats 80.41 % of golang submissions
// Your memory usage beats 87.95 % of golang submissions (7 MB)

// 链表
// 先计算两个链表的长度，差值为sub，长的那条先走sub步
// 然后两个链表一起走，直到相遇或走到最后为止

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lengthA, lengthB int
	p := headA
	for p != nil {
		p = p.Next
		lengthA++
	}
	p = headB
	for p != nil {
		p = p.Next
		lengthB++
	}
	if lengthA > lengthB {
		lengthA, lengthB = lengthB, lengthA
		headA, headB = headB, headA
	}

	pa, pb := headA, headB
	for i := 0; i < lengthB-lengthA; i++ {
		pb = pb.Next
	}

	for pa != pb && pa != nil {
		pa = pa.Next
		pb = pb.Next
	}

	return pa
}

// @lc code=end
