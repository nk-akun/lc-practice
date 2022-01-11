/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

package main

// @lc code=start

var ans []string
var ket []byte

func generateParenthesis(n int) []string {
	n *= 2
	ans = make([]string, 0)
	ket = make([]byte, n)

	dfs(0, 0, n)
	return ans
}

func dfs(pos, sum, n int) {
	if pos == n {
		if sum == 0 {
			ans = append(ans, string(ket))
		}
		return
	}

	if sum < (n - pos) {
		ket[pos] = '('
		dfs(pos+1, sum+1, n)
	}
	if sum > 0 {
		ket[pos] = ')'
		dfs(pos+1, sum-1, n)
	}
	return
}

// @lc code=end

// func main() {
// 	fmt.Println(generateParenthesis(3))
// }
