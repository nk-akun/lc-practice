/*
 * @lc app=leetcode.cn id=630 lang=golang
 *
 * [630] 课程表 III
 */

// 优先学习截止时间早的，发现时间不够时，放弃一门学过的课程中时间花费比当前的课程长的中的最长的 666
// 排序+优先队列
// 97/97 cases passed (116 ms)
// Your runtime beats 62.86 % of golang submissions
// Your memory usage beats 80 % of golang submissions (7.3 MB)

package main

import (
	"container/heap"
	"sort"
)

// @lc code=start
func scheduleCourse(courses [][]int) int {

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	ans := 0
	sum := 0
	c := &courseHeap{}
	heap.Init(c)
	for i := range courses {
		if sum+courses[i][0] > courses[i][1] {
			if c.Len() > 0 && (*c)[0] > courses[i][0] {
				m := heap.Pop(c)
				heap.Push(c, courses[i][0])
				sum = sum - m.(int) + courses[i][0]
			}
		} else {
			heap.Push(c, courses[i][0])
			sum += courses[i][0]
			ans++
		}
	}

	return ans
}

type courseHeap []int

func (c courseHeap) Len() int           { return len(c) }
func (c courseHeap) Less(i, j int) bool { return c[i] > c[j] }
func (c courseHeap) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c *courseHeap) Push(i interface{}) {
	*c = append(*c, i.(int))
}

func (c *courseHeap) Pop() interface{} {
	old := *c
	last := old[len(old)-1]
	*c = old[:len(old)-1]
	return last
}

// @lc code=end

// func main() {
// 	// fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3000}}))
// 	fmt.Println(scheduleCourse([][]int{{3, 3}, {2, 4}, {3, 5}}))
// }
