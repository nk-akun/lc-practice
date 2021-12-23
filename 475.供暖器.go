/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 */

// 直接二分答案，judge时判断每个房子是否有供暖器为其供暖，顺序判断houses时heaters也是顺序递增的，所以可以O(N)judge

// 执行用时：52 ms, 在所有 Go 提交中击败了29.17%的用户
// 内存消耗：6.6 MB, 在所有 Go 提交中击败了90.28%的用户
// 通过测试用例：30 / 30

package main

import (
	"sort"
)

// @lc code=start
func findRadius(houses []int, heaters []int) int {

	sort.Slice(houses, func(i, j int) bool { return houses[i] < houses[j] })
	sort.Slice(heaters, func(i, j int) bool { return heaters[i] < heaters[j] })

	l, r := 0, 1000000000
	for l <= r {
		mid := (l + r) >> 1
		if judge(houses, heaters, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func judge(houses []int, heaters []int, r int) bool {
	index := 0
	for _, h := range houses {
		minPos := h - r
		maxPos := h + r
		flag := false
		for index < len(heaters) && !flag {
			if heaters[index] >= minPos && heaters[index] <= maxPos {
				flag = true
			} else if heaters[index] < minPos {
				index++
			} else if heaters[index] > maxPos {
				return false
			}
		}
		if index == len(heaters) {
			return false
		}
	}

	return true
}

// @lc code=end

// func main() {
// 	// fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))
// 	// fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
// 	// fmt.Println(findRadius([]int{1, 5}, []int{2}))
// 	fmt.Println(findRadius([]int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}, []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}))
// }
