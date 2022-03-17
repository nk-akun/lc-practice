/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

//  70/70 cases passed (40 ms)
//  Your runtime beats 58.45 % of golang submissions
//  Your memory usage beats 40.4 % of golang submissions (9 MB)

// O(N)做法
// 将数字都记录到map中，然后从头遍历数组。
// ①针对遇到的数字x，在map中查找x+1、x+2、......直到遇到某个数在map中不存在为止，并标记走过的数，在map中查找x-1、x-2、......直到遇到某个数在map中不存在为止，并标记走过的数；
// ②继续遍历数组，如果下一个数已经被标记查过了，那么直接跳过，如果没查过，则执行①步骤；

package main

import (
	"fmt"
	"math"
)

// @lc code=start
func longestConsecutive(nums []int) int {
	existMap := make(map[int]bool)
	for i := range nums {
		existMap[nums[i]] = true
	}

	ans := 0
	for _, num := range nums {
		if v := existMap[num]; !v {
			continue
		}

		l, r := num, num
		for r = num + 1; r < math.MaxInt32; r++ {
			if _, exist := existMap[r]; !exist {
				r--
				break
			}
			existMap[r] = false
		}
		for l = num - 1; l > math.MinInt32; l-- {
			if _, exist := existMap[l]; !exist {
				l++
				break
			}
			existMap[l] = false
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
		existMap[num] = false
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1}))
}
