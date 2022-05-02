/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

package main

// 29/29 cases passed (36 ms)
// Your runtime beats 32.39 % of golang submissions
// Your memory usage beats 36.84 % of golang submissions (7.3 MB)

// 归并排序,合并两个有序链表
// 将链表从中间拆开，分别对两个子链表排序，然后将两个排好序的子链表归并即可

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	length := getLength(head)
	if length == 1 {
		return head
	}

	length /= 2
	p := head
	for i := 1; i < length; i++ {
		p = p.Next
	}

	var head1, head2 *ListNode
	head2 = sortList(p.Next)
	p.Next = nil
	head1 = sortList(head)

	return mergeList(head1, head2)
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	var head, it *ListNode
	p, q := head1, head2
	if p.Val > q.Val {
		head = q
		it = q
		q = q.Next
	} else {
		head = p
		it = p
		p = p.Next
	}

	for p != nil && q != nil {
		if p.Val < q.Val {
			it.Next = p
			p = p.Next
		} else {
			it.Next = q
			q = q.Next
		}
		it = it.Next
	}

	if p != nil {
		it.Next = p
	}

	if q != nil {
		it.Next = q
	}

	return head
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

// @lc code=end
