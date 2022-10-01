package main

import "fmt"

var n int
var m int
var last [][]int

func judgeLoop(mp []string, x, y int) (int, int) {
	// 判环
	vis := map[[2]int]bool{}
	for !vis[[2]int{x, y}] {
		vis[[2]int{x, y}] = true
		if last[x][y] > 0 {
			return -1, -1
		}
		switch mp[x][y] {
		case '^':
			x--
		case '>':
			y++
		case 'v':
			x++
		case '<':
			y--
		}
		if x < 0 || x >= n || y < 0 || y >= m {
			return -1, -1
		}
	}

	// 统计节点数
	count := 0
	vis = map[[2]int]bool{}
	for !vis[[2]int{x, y}] {
		vis[[2]int{x, y}] = true
		count++
		switch mp[x][y] {
		case '^':
			x--
		case '>':
			y++
		case 'v':
			x++
		case '<':
			y--
		}
	}

	// 环上节点赋值
	vis = map[[2]int]bool{}
	for !vis[[2]int{x, y}] {
		vis[[2]int{x, y}] = true
		last[x][y] = count
		switch mp[x][y] {
		case '^':
			x--
		case '>':
			y++
		case 'v':
			x++
		case '<':
			y--
		}
	}
	return x, y
}

func dfs(mp []string, x, y int) int {
	if x < 0 || x >= n || y < 0 || y >= m {
		return 0
	}
	if last[x][y] > 0 {
		return last[x][y]
	}
	switch mp[x][y] {
	case '^':
		x--
	case '>':
		y++
	case 'v':
		x++
	case '<':
		y--
	}
	return dfs(mp, x, y) + 1
}

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	mp := make([]string, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&mp[i])
	}

	last = make([][]int, n+1)
	for i := range last {
		last[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if last[i][j] != 0 {
				continue
			}
			x, _ := judgeLoop(mp, i, j)
			if x == -1 {
				last[i][j] = dfs(mp, i, j)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if last[i][j] > ans {
				ans = last[i][j]
			}
		}
	}

	fmt.Println(ans)
}

/*

2 3
^v<
>>^

4 4
v<<<
>>>v
^<<<
>>>>

*/
