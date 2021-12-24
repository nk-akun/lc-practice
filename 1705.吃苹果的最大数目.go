/*
 * @lc app=leetcode.cn id=1705 lang=golang
 *
 * [1705] 吃苹果的最大数目
 */

// 优先队列
// 把每天摘到的苹果数量以及最晚期限放到优先队列里
// 吃的时候不断从优先队列拿最早过期的，如果吃完了或者已经过期了就扔掉，重新拿

// 执行用时：184 ms, 在所有 Go 提交中击败了37.50%的用户
// 内存消耗：7.3 MB, 在所有 Go 提交中击败了37.50%的用户
// 通过测试用例：70 / 70

package main

import (
	"container/heap"
	"fmt"
)

// @lc code=start
func eatenApples(apples []int, days []int) int {
	aph := appleHeap{}
	heap.Init(&aph)

	var now node
	d, ans := 0, 0
	for {
		if d < len(apples) {
			heap.Push(&aph, node{count: apples[d], outTime: d + days[d] - 1})
		}
		if aph.Len() == 0 && d >= len(apples) {
			break
		}

		for aph.Len() > 0 {
			now = heap.Pop(&aph).(node)
			if now.count == 0 || now.outTime < d {
				continue
			}

			now.count--
			ans++
			heap.Push(&aph, now)
			break
		}
		d++
	}

	return ans
}

type node struct {
	count   int
	outTime int
}

type appleHeap []node

func (a appleHeap) Len() int           { return len(a) }
func (a appleHeap) Less(i, j int) bool { return a[i].outTime < a[j].outTime }
func (a appleHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a *appleHeap) Push(i interface{}) {
	*a = append(*a, i.(node))
}
func (a *appleHeap) Pop() interface{} {
	old := *a
	last := old[len(old)-1]
	*a = old[:len(old)-1]
	return last
}

// @lc code=end

func main() {
	fmt.Println(eatenApples([]int{100}, []int{10}))
}
