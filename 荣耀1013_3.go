package main

import "fmt"

var n, m int

func dfs(mp [][]int, sx, sy, ex, ey int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	dp[sx][sy] = 1
	qu := [][]int{}
	qu = append(qu, []int{sx, sy})

	ne := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(qu) > 0 {
		x, y := qu[0][0], qu[0][1]
		qu = qu[1:]

		for i := range ne {
			nx := x + ne[i][0]
			ny := y + ne[i][1]

			if nx < 0 || nx >= n || ny < 0 || ny >= n || mp[nx][ny] <= mp[x][y] {
				continue
			}
			dp[nx][ny] = (dp[nx][ny] + dp[x][y]) % 1000000000
			qu = append(qu, []int{nx, ny})
		}
	}
	return dp[ex][ey]
}

func main() {

	fmt.Scan(&n)
	fmt.Scan(&m)

	mp := make([][]int, n)
	for i := range mp {
		mp[i] = make([]int, m)
	}

	for i := range mp {
		for j := range mp[i] {
			fmt.Scan(&mp[i][j])
		}
	}

	var sx, sy, ex, ey int
	fmt.Scan(&sx)
	fmt.Scan(&sy)
	fmt.Scan(&ex)
	fmt.Scan(&ey)

	fmt.Println(dfs(mp, sx, sy, ex, ey))
}

/*

4 5
0 1 0 0 0
0 2 3 0 0
0 3 4 5 0
0 0 7 6 0
0 1 2 3

*/
