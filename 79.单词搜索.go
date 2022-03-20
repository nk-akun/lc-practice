/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

//  83/83 cases passed (92 ms)
//  Your runtime beats 56.25 % of golang submissions
//  Your memory usage beats 49.57 % of golang submissions (1.9 MB)

package main

import "fmt"

// @lc code=start

// 写崩了，不应该用BFS
// func dfs(x int, y int, word string, board [][]byte, dp [][][]bool) bool {
// 	if board[x][y] != word[0] {
// 		return false
// 	}
// 	ne := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
// 	vis := make([][]bool, len(board))
// 	for i := range vis {
// 		vis[i] = make([]bool, len(board[0]))
// 	}
// 	queue := make([][]int, 0)
// 	queue = append(queue, []int{x, y, 0})
// 	l := -1

// 	for l < len(queue)-1 {
// 		l++

// 		now := queue[l]
// 		if dp[now[0]][now[1]][now[2]] {
// 			continue
// 		}
// 		dp[now[0]][now[1]][now[2]] = true
// 		vis[now[0]][now[1]] = true
// 		if now[2] == len(word)-1 {
// 			return true
// 		}
// 		for i := range ne {
// 			nx := now[0] + ne[i][0]
// 			ny := now[1] + ne[i][1]
// 			if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[0]) || board[nx][ny] != word[now[2]+1] || vis[nx][ny] {
// 				continue
// 			}

// 			// if dp[nx][ny][now[2]+1] {
// 			// 	// 待优化
// 			// 	continue
// 			// }
// 			queue = append(queue, []int{nx, ny, now[2] + 1})
// 		}
// 	}

// 	return false
// }

// func exist(board [][]byte, word string) bool {
// 	dp := make([][][]bool, len(board))
// 	for i := range dp {
// 		dp[i] = make([][]bool, len(board[i]))
// 		for j := range dp[i] {
// 			dp[i][j] = make([]bool, len(word))
// 		}
// 	}

// 	flag := false
// 	for i := range board {
// 		for j := range board[i] {
// 			if dfs(i, j, word, board, dp) {
// 				flag = true
// 				break
// 			}
// 		}
// 		if flag {
// 			break
// 		}
// 	}
// 	return flag
// }

func dfs(x int, y int, word string, board [][]byte, vis [][]bool, pos int) bool {
	if pos == len(word)-1 {
		return true
	}

	vis[x][y] = true
	ne := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := range ne {
		nx := x + ne[i][0]
		ny := y + ne[i][1]
		if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[0]) || board[nx][ny] != word[pos+1] || vis[nx][ny] {
			continue
		}
		if dfs(nx, ny, word, board, vis, pos+1) {
			return true
		}
	}
	vis[x][y] = false
	return false
}

func exist(board [][]byte, word string) bool {

	flag := false
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				vis := make([][]bool, len(board))
				for i := range vis {
					vis[i] = make([]bool, len(board[0]))
				}
				if dfs(i, j, word, board, vis, 0) {
					flag = true
					break
				}
			}
		}
		if flag {
			break
		}
	}
	return flag
}

// @lc code=end

func main() {
	// fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'a', 'a'}}, "aaa"))
}
