/*
 * @lc app=leetcode.cn id=488 lang=golang
 *
 * [488] 祖玛游戏
 */

package main

import "fmt"

// @lc code=start

var ans int = 999

func findMinStep(board string, hand string) int {
	mp := map[byte]int{}
	for i := range hand {
		mp[hand[i]]++
	}

	for i := range board {
		if i > 0 && board[i] == board[i-1] {
			continue
		}
		dfs(mp, board, i, 0)
	}

	if ans == 999 {
		return -1
	}
	return ans
}

func dfs(mp map[byte]int, board string, i int, cost int) {
	// fmt.Println(board, mp, i, cost)
	b := board[i]
	var j int
	for j = i + 1; j < len(board) && board[j] == board[j-1]; j++ {
	}
	c := 3 - (j - i)
	if mp[b] < c {
		return
	}
	mp[b] -= c
	cost += c
	if cost >= ans {
		return
	}
	// fmt.Println(board)
	board = board[:i] + board[j:]
	// fmt.Println(board)
	board = remove(board, i-1)
	if board == "" {
		// fmt.Println(mp)
		if cost < ans {
			ans = cost
		}
		return
	}

	for k := range board {
		if k > 0 && board[k] == board[k-1] {
			continue
		}
		dfs(mp, board, k, cost)
	}
	mp[b] += c
	return
}

func remove(s string, i int) string {
	if i < 0 || s == "" {
		return s
	}
	var j int
	for j = i + 1; j < len(s) && s[j] == s[j-1]; j++ {
	}
	for i = i - 1; i >= 0 && s[i] == s[i+1]; i-- {
	}
	if j-i-1 < 3 {
		return s
	}
	return remove(s[:i+1]+s[j:], i)
}

// @lc code=end

func main() {
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))
}

// "RBYYBBRRB"
// "YRBGB"
