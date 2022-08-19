/*
 * @lc app=leetcode.cn id=1171 lang=golang
 *
 * [1171] 从链表中删去总和值为零的连续节点
 */

package main

// 105/105 cases passed (4 ms)
// Your runtime beats 87.88 % of golang submissions
// Your memory usage beats 45.45 % of golang submissions (3.7 MB)

// 前缀和
// 需要把每个节点的val设置为前缀和，并用map记录每个前缀和是否出现过，如果出现过，再这个区间加和为0
// 最后再遍历一遍把其从前缀和改为本来的值

import "fmt"

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func deletenode(head *ListNode, start *ListNode, end *ListNode) {
	p := head
	for p.Next != start {
		p = p.Next
	}

	p.Next = end.Next
	return
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	newHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	vis := map[int]*ListNode{}
	vis[0] = newHead
	last := 0
	p := newHead.Next
	for p != nil {
		last += p.Val
		p.Val = last

		if vis[last] != nil {
			start := vis[last].Next
			end := p

			for tmp := start; tmp != end; tmp = tmp.Next {
				vis[tmp.Val] = nil
			}
			p = p.Next
			deletenode(newHead, start, end)
		} else {
			vis[last] = p
			p = p.Next
		}
	}

	last = 0
	for p := newHead.Next; p != nil; p = p.Next {
		tmp := p.Val
		p.Val -= last
		last = tmp
	}

	return newHead.Next
}

// @lc code=end

func main() {
	// head := &ListNode{1, &ListNode{2, &ListNode{-3, &ListNode{3, &ListNode{1, nil}}}}}
	// head = removeZeroSumSublists(head)
	// for head != nil {
	// 	fmt.Println(head.Val)
	// 	head = head.Next
	// }

	head := &ListNode{0, nil}
	head = removeZeroSumSublists(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
