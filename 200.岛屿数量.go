/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

package main

// 49/49 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 23.56 % of golang submissions (4 MB)

// dfs搞一搞
// 深搜把走过的点都标记一下，所以只要没走过的都是新的一块岛屿

// @lc code=start
func numIslands(grid [][]byte) int {
	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[0]))
	}

	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !vis[i][j] {
				ans++
				dfs(i, j, vis, grid)
			}
		}
	}
	return ans
}

func dfs(x, y int, vis [][]bool, grid [][]byte) {
	ne := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	vis[x][y] = true

	for i := range ne {
		nx := x + ne[i][0]
		ny := y + ne[i][1]
		if nx < 0 || nx >= len(vis) || ny < 0 || ny >= len(vis[0]) || vis[nx][ny] || grid[nx][ny] == '0' {
			continue
		}
		dfs(nx, ny, vis, grid)
	}
	return
}

// @lc code=end
