package main

import (
	"container/heap"
)

func solvesf(n int, m int, a []int) int {
	h := intHeap{}
	heap.Init(&h)

	for i := range a {
		heap.Push(&h, a[i])
	}

	ans := 0
	topM := make([]int, m)
	for h.Len() >= m {
		for i := range topM {
			topM[i] = heap.Pop(&h).(int)
		}
		now := topM[len(topM)-1]
		ans += now
		for i := range topM {
			if topM[i] > now {
				heap.Push(&h, topM[i]-now)
			}
		}
	}

	return ans
}

type intHeap []int

func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
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
