/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// 深搜的方式进行匹配

package main

// @lc code=start
func isMatch(s string, p string) bool {
	if dfs(s, p) {
		return true
	}
	return false
}

func dfs(s string, p string) bool {

	// 特判
	if s != "" && p == "" {
		return false
	} else if s == "" && p == "" {
		return true
	}

	// 如果第二位是*
	if len(p) >= 2 && p[1] == '*' {
		// 第一位是.的话，可以匹配任何字符
		if p[0] == '.' {
			for i := 0; i <= len(s); i++ {
				if dfs(s[i:], p[2:]) {
					return true
				}
				if i >= len(s) {
					break
				}
			}
		}
		// 如果首位相同，for遍历s首位取多少个
		if len(s) > 0 && s[0] == p[0] {
			for i := 0; i <= len(s); i++ {
				if dfs(s[i:], p[2:]) {
					return true
				}
				if i >= len(s) || (i > 0 && s[i] != s[i-1]) {
					break
				}
			}
		} else if dfs(s, p[2:]) {
			// 首位不同，直接p[0]取0个
			return true
		}
		return false
	} else {
		// 第二位不是*，首位相同或不同都可以继续向下匹配
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			if dfs(s[1:], p[1:]) {
				return true
			}
		}
		return false
	}
}

// @lc code=end

// func main() {
// 	fmt.Println(isMatch("ab", ".*"))
// }
