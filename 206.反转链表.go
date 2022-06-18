/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

package main

// 28/28 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 99.96 % of golang submissions (2.4 MB)

// 记得 头结点.Next = nil,否则前两个节点出环了

import "fmt"

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p, q := head, head.Next
	head.Next = nil
	for q != nil {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}

	return p
}

// @lc code=end

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = reverseList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
