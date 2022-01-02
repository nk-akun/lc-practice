/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 */

// 思维题
// 数字数量每次除以二，数字之间的间隔每次乘以2，所以只需要记录最小值、最大值即可
// 每轮之后，动态调整最小值、最大值，二者相等时结束

package main

// @lc code=start
func lastRemaining(n int) int {
	minNum, maxNum := 1, n
	i := 1
	remainLen := n
	for minNum != maxNum {
		if i&1 == 1 {
			if remainLen&1 == 1 {
				minNum = minNum + (1 << (i - 1))
				maxNum = maxNum - (1 << (i - 1))
			} else {
				minNum = minNum + (1 << (i - 1))
			}
		} else {
			if remainLen&1 == 1 {
				minNum = minNum + (1 << (i - 1))
				maxNum = maxNum - (1 << (i - 1))
			} else {
				maxNum = maxNum - (1 << (i - 1))
			}
		}
		i++
		remainLen >>= 1
	}
	return minNum
}

// @lc code=end

// func main() {
// 	fmt.Println(lastRemaining(1000000000))
// }
