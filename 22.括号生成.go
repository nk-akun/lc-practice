/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// 深搜+剪枝

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

	// 如果剩余的都是有括号，能否抵掉当前的左括号
	if sum < (n - pos) {
		ket[pos] = '('
		dfs(pos+1, sum+1, n)
	}

	// 只有有左括号时，才能有有括号
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
