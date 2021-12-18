/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */

// O(N)扫描，对于每个X看看有没有上或左与之相连的，没有则ans++

//  执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//  内存消耗：2.3 MB, 在所有 Go 提交中击败了21.43%的用户
//  通过测试用例：27 / 27

package main

// @lc code=start
func countBattleships(board [][]byte) int {
	ans := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X' {
					continue
				}
				ans++
			}
		}
	}
	return ans
}

// @lc code=end
