/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// 先让front指针走n-1步，behind指针指向表头。然后front和behind一起走，直到front到表尾，删除behind指向的节点。
// 这里注意如果删除的是表头需要特殊处理

//  208/208 cases passed (0 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 62.38 % of golang submissions (2.2 MB)

package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	front, behind := head, head
	var last *ListNode = nil

	for n > 1 {
		front = front.Next
		n--
	}

	for front.Next != nil {
		last = behind
		front = front.Next
		behind = behind.Next
	}

	var ans *ListNode = nil
	if behind == head {
		ans = behind.Next
		behind.Next = nil
	} else {
		ans = head
		last.Next = behind.Next
		behind.Next = nil
	}

	return ans
}

// @lc code=end
