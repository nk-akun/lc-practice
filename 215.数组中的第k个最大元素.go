/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

package main

// 32/32 cases passed (4 ms)
// Your runtime beats 97.24 % of golang submissions
// Your memory usage beats 10.49 % of golang submissions (4 MB)

// k大的优先队列跑一跑

import (
	"container/heap"
	"fmt"
)

// @lc code=start
func findKthLargest(nums []int, k int) int {
	h := intHeap{}
	heap.Init(&h)
	for i := range nums {
		if h.Len() < k {
			heap.Push(&h, nums[i])
		} else {
			top := heap.Pop(&h).(int)
			if top < nums[i] {
				heap.Push(&h, nums[i])
			} else {
				heap.Push(&h, top)
			}
		}
	}
	return heap.Pop(&h).(int)
}

type intHeap []int

func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Len() int           { return len(h) }

func (h *intHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	last := old[len(old)-1]
	*h = old[:len(old)-1]
	return last
}

// @lc code=end

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 1))
}
