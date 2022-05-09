/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

package main

// 21/21 cases passed (8 ms)
// Your runtime beats 98.16 % of golang submissions
// Your memory usage beats 19.67 % of golang submissions (5.4 MB)

// 优先队列
// map计数+优先队列

import (
	"container/heap"
	"fmt"
)

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for i := range nums {
		mp[nums[i]]++
	}

	h := numHeap{}
	heap.Init(&h)
	for key, v := range mp {
		if h.Len() < k {
			heap.Push(&h, []int{key, v})
		} else {
			top := heap.Pop(&h).([]int)
			if v > top[1] {
				heap.Push(&h, []int{key, v})
			} else {
				heap.Push(&h, top)
			}
		}
	}

	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(&h).([]int)[0])
	}
	return ans
}

type numHeap [][]int

func (h numHeap) Len() int           { return len(h) }
func (h numHeap) Swap(x, y int)      { h[x], h[y] = h[y], h[x] }
func (h numHeap) Less(x, y int) bool { return h[x][1] < h[y][1] }

func (h *numHeap) Push(i interface{}) {
	*h = append(*h, i.([]int))
}

func (h *numHeap) Pop() interface{} {
	old := *h
	result := old[len(old)-1]
	*h = old[:len(old)-1]
	return result
}

// @lc code=end

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 4, 4, 5}, 3))
}
