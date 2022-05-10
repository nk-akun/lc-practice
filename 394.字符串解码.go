/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

package main

// 34/34 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 14.16 % of golang submissions (2.2 MB)

// 递归，时间O(N)
// 遇到字母直接append,遇到数字就把数字择出来然后递归处理数字后面[]内的string

import "fmt"

// @lc code=start
func decodeString(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			cnt := 0
			for ; i < len(s) && (s[i] >= '0' && s[i] <= '9'); i++ {
				cnt = cnt*10 + int(s[i]-'0')
			}

			i++
			r := 1
			var j int
			for j = i; j < len(s) && r > 0; j++ {
				if s[j] == ']' {
					r--
				} else if s[j] == '[' {
					r++
				}
			}
			j--

			t := decodeString(s[i:j])
			for cnt > 0 {
				result += t
				cnt--
			}
			i = j
		} else {
			result += string(s[i])
		}
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
	fmt.Println(decodeString("aaaaaaa4[b3[c2[a1[d]]]]"))
}
