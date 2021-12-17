/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// dfs暴力乱搞

//  25/25 cases passed (0 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 99 % of golang submissions (2 MB)

package main

// @lc code=start
var ans []string

func letterCombinations(digits string) []string {
	ans = []string{}
	if len(digits) == 0 {
		return ans
	}
	mp := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	dfs(digits, mp, 0, []byte{})

	return ans
}

func dfs(digits string, mp map[byte]string, pos int, current []byte) {
	if pos >= len(digits) {
		ans = append(ans, string(current))
		return
	}

	now := digits[pos]
	for _, b := range mp[now] {
		current = append(current, byte(b))
		dfs(digits, mp, pos+1, current)
		current = current[:len(current)-1]
	}
	return
}

// @lc code=end

// func main() {
// 	fmt.Println(letterCombinations(""))
// }
