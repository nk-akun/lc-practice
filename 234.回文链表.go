/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

package main

// 86/86 cases passed (112 ms)
// Your runtime beats 94.13 % of golang submissions
// Your memory usage beats 24.06 % of golang submissions (11 MB)

// 反转链表
// 通过把后半部分翻转过来，比较前半部和后半部分即可

// 或者用快慢指针的方式，一个指针每次走一步，另一个指针每次走两步，这样第二个指针走到最后时，第一个指针正好在最中间

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func isPalindrome(head *ListNode) bool {
	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}

	if length == 1 {
		return true
	}

	p = head
	for i := 1; i < length/2; i++ {
		p = p.Next
	}
	if (length & 1) == 1 {
		p = p.Next
	}
	for q := p.Next; q != nil; {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}

	t := head
	for i := 1; i <= length/2; i++ {
		if t.Val != p.Val {
			return false
		}
		t = t.Next
		p = p.Next
	}

	return true
}

// @lc code=end
