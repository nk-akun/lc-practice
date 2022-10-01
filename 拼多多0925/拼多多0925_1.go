package main

import "fmt"

var n, m int         // 节点数，猫粮数
var a []int          // 猫数
var vis map[int]bool // 标记有没有走过
var mp [][]int       // 地图

var ans int     // 走过的最多点
var minCost int // 最小猫粮耗费

func dfs(x int, count int, cost int) {
	if count > ans {
		ans = count
		minCost = m - cost
	} else if count == ans {
		if m-cost < minCost {
			minCost = m - cost
		}
	}
	vis[x] = true

	for _, nt := range mp[x] {
		if vis[nt] || a[nt] > cost {
			continue
		}
		dfs(nt, count+1, cost-a[nt])
	}
}

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	vis = make(map[int]bool)
	a = make([]int, n+1)
	mp = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	var x, y int
	for i := 1; i < n; i++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		if len(mp[x]) == 0 {
			mp[x] = make([]int, 0)
		}
		mp[x] = append(mp[x], y)
		if len(mp[y]) == 0 {
			mp[y] = make([]int, 0)
		}
		mp[y] = append(mp[y], x)
	}

	if m >= a[1] {
		dfs(1, 1, m-a[1])
	}
	fmt.Printf("%d %d\n", ans, minCost)
}
