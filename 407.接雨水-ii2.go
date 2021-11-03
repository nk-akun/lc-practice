/*
 * @lc app=leetcode.cn id=407 lang=golang
 *
 * [407] 接雨水 II
 */

package main

import (
	"container/heap"
)

// 先将边界节点push到优先队列中，优先队列以height为优先级，低的在堆头
// 然后向内部扩散，可以想到：当前pop出来的点，如果周围的点比其低且没有压入过队列，那么该点一点可以接水
// 本着以上原则就可以解决此题

// @lc code=start
func trapRainWater(heightMap [][]int) int {
	pQueue := &hp{}

	vis := make([][]bool, len(heightMap))
	for i := 0; i < len(vis); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			vis[i] = append(vis[i], false)
		}
	}
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if i == 0 || j == 0 || i == len(heightMap)-1 || j == len(heightMap[0])-1 {
				heap.Push(pQueue, cell{heightMap[i][j], i, j})
				vis[i][j] = true
			}
		}
	}

	ans := 0
	ne := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for pQueue.Len() > 0 {
		now := heap.Pop(pQueue).(cell)
		for i := range ne {
			nx := now.x + ne[i][0]
			ny := now.y + ne[i][1]
			if nx < 0 || nx >= len(heightMap) || ny < 0 || ny >= len(heightMap[0]) || vis[nx][ny] {
				continue
			}
			if now.h > heightMap[nx][ny] {
				ans += now.h - heightMap[nx][ny]
			}
			heap.Push(pQueue, cell{max(heightMap[nx][ny], now.h), nx, ny})
			vis[nx][ny] = true
		}
	}

	return ans
}

type cell struct{ h, x, y int }
type hp []cell

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(cell)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

// func main() {

// 	// heightMap := [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}
// 	// heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
// 	// heightMap := [][]int{{12, 13, 1, 12}, {13, 4, 13, 12}, {13, 8, 10, 12}, {12, 13, 12, 12}, {13, 13, 13, 13}}
// 	// heightMap := [][]int{{5, 5, 5, 1}, {5, 1, 1, 5}, {5, 1, 5, 5}, {5, 2, 5, 8}}
// 	// heightMap := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
// 	// heightMap := [][]int{{78, 16, 94, 36}, {87, 93, 50, 22}, {63, 28, 91, 60}, {64, 27, 41, 27}, {73, 37, 12, 69}, {68, 30, 83, 31}, {63, 24, 68, 36}}
// 	fmt.Println(trapRainWater(heightMap))
// }
