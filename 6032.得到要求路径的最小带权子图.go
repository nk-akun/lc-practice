package main

// 执行用时：320 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：33.6 MB, 在所有 Go 提交中击败了100.00%的用户
// 通过测试用例：78 / 78

// 我们的直接想法是：让src1和src2通过最短路算法分别走向dest，那么连接成的子图是最小的(一定可以连接成子图，因为终点相同)
// 但是我们不知道他俩路线的公共部分是哪些
// 所以我们的做法就是枚举相遇点，计算src1到相遇点和src2到相遇点的最短路，再加上相遇点到dest的最短路
// 相遇点到dest的最短路可以通过求dest到所有点的最短路得出

// Dijkstra

import (
	"container/heap"
	"fmt"
)

func dj(edges [][][]int, start int) []int64 {
	dis := make([]int64, len(edges)+10)
	for i := range dis {
		dis[i] = 1e11
	}
	dis[start] = 0

	hp := destHeap{}
	heap.Init(&hp)
	heap.Push(&hp, node{start, 0})

	for hp.Len() > 0 {
		now := heap.Pop(&hp).(node)
		if dis[now.v] < now.dis {
			continue
		}
		for _, e := range edges[now.v] {

			next := e[0]
			cost := e[1]
			if dis[now.v]+int64(cost) < dis[next] {
				dis[next] = dis[now.v] + int64(cost)
				heap.Push(&hp, node{next, dis[next]})
			}
		}
	}
	return dis
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	eds := make([][][]int, n)
	edrs := make([][][]int, n)
	for _, e := range edges {
		eds[e[0]] = append(eds[e[0]], []int{e[1], e[2]})
		edrs[e[1]] = append(edrs[e[1]], []int{e[0], e[2]})
	}

	dis1 := dj(eds, src1)
	dis2 := dj(eds, src2)
	disr := dj(edrs, dest)

	var ans int64 = 1e11
	for i := 0; i < n; i++ {
		if dis1[i]+dis2[i]+disr[i] < ans {
			ans = dis1[i] + dis2[i] + disr[i]
		}
	}

	if ans == 1e11 {
		return -1
	}
	return ans
}

type node struct {
	v   int
	dis int64
}
type destHeap []node

func (d destHeap) Len() int           { return len(d) }
func (d destHeap) Less(i, j int) bool { return d[i].dis < d[j].dis }
func (d destHeap) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d *destHeap) Push(i interface{}) {
	*d = append(*d, i.(node))
}
func (d *destHeap) Pop() interface{} {
	old := *d
	result := old[len(old)-1]
	*d = old[:len(old)-1]
	return result
}

func main() {
	fmt.Println(minimumWeight(6, [][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}}, 0, 1, 5))
}
