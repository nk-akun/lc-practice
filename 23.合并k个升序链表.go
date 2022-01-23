/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// 133/133 cases passed (8 ms)
// Your runtime beats 91.68 % of golang submissions
// Your memory usage beats 38.89 % of golang submissions (5.6 MB)

// 优先队列

package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func mergeKLists(lists []*ListNode) *ListNode {
	h := listHeap{}
	heap.Init(&h)
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		heap.Push(&h, lists[i])
	}
	if h.Len() == 0 {
		return nil
	}

	ans := heap.Pop(&h).(*ListNode)
	if ans.Next != nil {
		heap.Push(&h, ans.Next)
	}
	last := ans

	for h.Len() > 0 {
		now := heap.Pop(&h).(*ListNode)
		last.Next = now
		last = now
		if now.Next != nil {
			heap.Push(&h, now.Next)
		}
	}
	return ans
}

type listHeap []*ListNode

func (l listHeap) Less(i, j int) bool { return l[i].Val < l[j].Val }
func (l listHeap) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l listHeap) Len() int           { return len(l) }

func (l *listHeap) Push(i interface{}) {
	*l = append(*l, i.(*ListNode))
}
func (l *listHeap) Pop() interface{} {
	old := *l
	last := old[len(old)-1]
	*l = old[:len(old)-1]
	return last
}

// @lc code=end
