package main

import (
	"container/heap"
	"fmt"
)

var n, m int
var mp []string

func bfs() int {
	h := mHeap{}
	heap.Init(&h)

	vis := make([][][]bool, 2)
	for i := 0; i < n; i++ {
		vis[0] = append(vis[0], make([]bool, m))
	}
	for i := 0; i < n; i++ {
		vis[1] = append(vis[1], make([]bool, m))
	}

	ne := [][][]int{}
	ne = append(ne, [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}})
	ne = append(ne, [][]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}})
	heap.Push(&h, node{
		x:    0,
		y:    0,
		tp:   0,
		cost: 0,
	})

	for h.Len() > 0 {
		now := heap.Pop(&h).(node)
		if vis[now.tp][now.x][now.y] {
			continue
		}
		vis[now.tp][now.x][now.y] = true
		if now.x == n-1 && now.y == m-1 {
			return now.cost
		}

		for i := range ne[now.tp] {
			nx := now.x + ne[now.tp][i][0]
			ny := now.y + ne[now.tp][i][1]
			if nx < 0 || ny < 0 || nx >= n || ny >= m || vis[now.tp][nx][ny] || mp[nx][ny] == 'X' {
				continue
			}
			heap.Push(&h, node{
				x:    nx,
				y:    ny,
				tp:   now.tp,
				cost: now.cost + 1,
			})
		}

		if mp[now.x][now.y] == 'S' && !vis[now.tp^1][now.x][now.y] {
			heap.Push(&h, node{
				x:    now.x,
				y:    now.y,
				tp:   now.tp ^ 1,
				cost: now.cost + 1,
			})
		}
	}
	return -1
}

func main() {

	fmt.Scan(&m)
	fmt.Scan(&n)

	mp = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mp[i])
	}

	fmt.Println(bfs())
}

type node struct {
	x    int
	y    int
	tp   int
	cost int
}

type mHeap []node

func (a mHeap) Len() int           { return len(a) }
func (a mHeap) Less(i, j int) bool { return a[i].cost < a[j].cost }
func (a mHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a *mHeap) Push(i interface{}) {
	*a = append(*a, i.(node))
}
func (a *mHeap) Pop() interface{} {
	old := *a
	last := old[len(old)-1]
	*a = old[:len(old)-1]
	return last
}
