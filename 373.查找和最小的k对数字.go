/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// 执行用时：80 ms, 在s所有 Go 提交中击败了96.05%的用户
// 内存消耗：9.6 MB, 在所有 Go 提交中击败了68.42%的用户
// 通过测试用例：32 / 32

// 优先队列
// 通过优先队列找出最小的组合来
// 加入优先队列的条件是：
// 1.pop出来一个组合，将(pos1,pos2+1)组合push进去
// 2.pop出来一个组合，如果pos2==0，将(pos1+1,0)组合push进去
// 这样可以保证push进优先队列的组合尽量少，节省内存

package main

import (
	"container/heap"
	"fmt"
)

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	count := 0
	shp := sumHeap{}
	heap.Init(&shp)
	heap.Push(&shp, node{0, 0, nums1[0] + nums2[0]})
	ans := make([][]int, 0, k)

	var now node
	for count < k && shp.Len() > 0 {
		now = heap.Pop(&shp).(node)
		ans = append(ans, []int{nums1[now.pos1], nums2[now.pos2]})
		if now.pos2 < len(nums2)-1 {
			heap.Push(&shp, node{now.pos1, now.pos2 + 1, nums1[now.pos1] + nums2[now.pos2+1]})
		}
		if now.pos2 == 0 && now.pos1 < len(nums1)-1 {
			heap.Push(&shp, node{now.pos1 + 1, 0, nums1[now.pos1+1] + nums2[0]})
		}
		count++
	}

	return ans
}

type node struct {
	pos1 int
	pos2 int
	sum  int
}
type sumHeap []node

func (s sumHeap) Len() int           { return len(s) }
func (s sumHeap) Less(i, j int) bool { return s[i].sum < s[j].sum }
func (s sumHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *sumHeap) Push(i interface{}) {
	*s = append(*s, i.(node))
}
func (s *sumHeap) Pop() interface{} {
	old := *s
	last := old[len(old)-1]
	*s = old[:len(old)-1]
	return last
}

// @lc code=end

func main() {
	// fmt.Println(kSmallestPairs([]int{1, 4, 5, 11, 13, 17}, []int{1, 3, 4, 6, 11, 18}, 100))
	fmt.Println(kSmallestPairs([]int{1, 2}, []int{3}, 3))
}
