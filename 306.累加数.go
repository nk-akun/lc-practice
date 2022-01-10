/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// 双重for循环先把前两个数确定，后面的数都是可计算的
// 注意0的判断即可

package main

import (
	"strconv"
)

// @lc code=start
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	for i := 0; i < len(num)-2; i++ {
		for j := i + 1; j < len(num)-1; j++ {
			num1, num2 := 0, 0
			for k := 0; k <= i; k++ {
				num1 = num1*10 + int(num[k]-'0')
			}
			for k := i + 1; k <= j; k++ {
				num2 = num2*10 + int(num[k]-'0')
			}
			nums := []int{num1, num2}
			if judge(nums, num, j+1) {
				return true
			}

			// 保证第二个数只有一个0
			if num[i+1] == '0' {
				break
			}
		}
		// 保证第一个数只有一个0
		if num[0] == '0' {
			break
		}
	}
	return false
}

func judge(nums []int, s string, pos int) bool {
	if pos == len(s) {
		return true
	}
	sum := nums[0] + nums[1]
	sumS := strconv.Itoa(sum)
	for i := 0; i < len(sumS); i++ {
		if i+pos >= len(s) || sumS[i] != s[i+pos] {
			return false
		}
	}
	nums[0] = nums[1]
	nums[1] = sum

	return judge(nums, s, pos+len(sumS))
}

// @lc code=end

// func main() {
// 	fmt.Println(isAdditiveNumber("10112"))
// }
