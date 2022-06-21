/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

package main

// 62/62 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 75.89 % of golang submissions (3.4 MB)

import (
	"fmt"
)

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	var last *ListNode
	p := head
	for p != nil {
		h, t := p, p
		for i := 1; i < k && t != nil; i++ {
			t = t.Next
		}
		if t == nil {
			break
		}

		p = t.Next
		h, t = reverse(h, t)
		if last != nil {
			last.Next = h
		} else {
			head = h
		}
		t.Next = p
		last = t
	}

	return head
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	p, q := head, head.Next
	for p != tail {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}

	return tail, head
}

// @lc code=end

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	head = reverseKGroup(head, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
