/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

package main

import "fmt"

// @lc code=start
func getHint(secret string, guess string) string {
	a, b := 0, 0
	count := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			count[secret[i]-'0']++
		}
	}
	for i := 0; i < len(secret); i++ {
		if secret[i] != guess[i] {
			if count[guess[i]-'0'] > 0 {
				b++
				count[guess[i]-'0']--
			}
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}

// @lc code=end

// func main() {
// 	fmt.Println(getHint("111087", "100781"))
// }
