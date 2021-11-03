/*
 * @lc app=leetcode.cn id=407 lang=golang
 *
 * [407] 接雨水 II
 */
package main

// 解法1:先预处理每个点的纵列和横列的最大值，表示二维下的最大可接雨水高度，但如果周围点可接雨水高度低于当前点，雨水仍然可以通过周围点流走，
// 所以需要用bfs的方式，从任意一个点往外扩散，如果周围点比当前点低，那么周围点可接雨水高度就需要降低,直到没有任何一个点可以降低为止

// 解法2:先假设除边界点的所有节点可接雨水高度都等同于矩阵中的最高点
// 然后用bfs的方式，从任意一个点往外扩散，如果周围点比当前点低，那么周围点可接雨水高度就需要降低，直到没有任何一个点可以降低为止

// @lc code=start
func trapRainWater(heightMap [][]int) int {
	hi := make([][]int, len(heightMap))
	for i := 0; i < len(heightMap); i++ {
		hi[i] = make([]int, len(heightMap[0]))
	}

	// 解法1
	// for i := 0; i < len(heightMap); i++ {
	// 	maxh := 0
	// 	for j := 0; j < len(heightMap[i]); j++ {
	// 		maxh = max(heightMap[i][j], maxh)
	// 		hi[i][j] = maxh
	// 	}
	// }

	// for i := 0; i < len(heightMap); i++ {
	// 	maxh := 0
	// 	for j := len(heightMap[i]) - 1; j >= 0; j-- {
	// 		maxh = max(heightMap[i][j], maxh)
	// 		hi[i][j] = min(hi[i][j], maxh)
	// 	}
	// }

	// for i := 0; i < len(heightMap[0]); i++ {
	// 	maxh := 0
	// 	for j := 0; j < len(heightMap); j++ {
	// 		maxh = max(heightMap[j][i], maxh)
	// 		hi[j][i] = min(hi[j][i], maxh)
	// 	}
	// }

	// for i := 0; i < len(heightMap[0]); i++ {
	// 	maxh := 0
	// 	for j := len(heightMap) - 1; j >= 0; j-- {
	// 		maxh = max(heightMap[j][i], maxh)
	// 		hi[j][i] = min(hi[j][i], maxh)
	// 	}
	// }

	// 解法2
	maxH := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			maxH = max(maxH, heightMap[i][j])
		}
	}

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			if i == 0 || j == 0 || i == len(heightMap)-1 || j == len(heightMap[0])-1 {
				hi[i][j] = heightMap[i][j]
			} else {
				hi[i][j] = maxH
			}
		}
	}

	// 两个解法的公用部分
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(hi); i++ {
			for j := 0; j < len(hi[i]); j++ {
				if num := bfs(hi, heightMap, i, j); num > 0 {
					flag = true
				}
			}
		}
	}

	ans := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if hi[i][j] > heightMap[i][j] {
				ans += hi[i][j] - heightMap[i][j]
			}
		}
	}

	return ans
}

func bfs(hi [][]int, heightMap [][]int, sx, sy int) int {
	result := 0

	type node struct{ x, y int }
	ne := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := []node{}
	queue = append(queue, node{x: sx, y: sy})
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		for i := range ne {
			nx := now.x + ne[i][0]
			ny := now.y + ne[i][1]
			if nx < 0 || nx >= len(hi) || ny < 0 || ny >= len(hi[0]) || hi[nx][ny] == heightMap[nx][ny] || hi[now.x][now.y] >= hi[nx][ny] {
				continue
			}
			result++
			hi[nx][ny] = max(heightMap[nx][ny], hi[now.x][now.y])
			queue = append(queue, node{x: nx, y: ny})
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

// func main() {

// 	// heightMap := [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}
// 	// heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
// 	// heightMap := [][]int{{12, 13, 1, 12}, {13, 4, 13, 12}, {13, 8, 10, 12}, {12, 13, 12, 12}, {13, 13, 13, 13}}
// 	// heightMap := [][]int{{5, 5, 5, 1}, {5, 1, 1, 5}, {5, 1, 5, 5}, {5, 2, 5, 8}}
// 	// heightMap := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
// 	heightMap := [][]int{{78, 16, 94, 36}, {87, 93, 50, 22}, {63, 28, 91, 60}, {64, 27, 41, 27}, {73, 37, 12, 69}, {68, 30, 83, 31}, {63, 24, 68, 36}}
// 	fmt.Println(trapRainWater(heightMap))
// }
