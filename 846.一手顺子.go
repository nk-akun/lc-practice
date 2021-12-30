/*
 * @lc app=leetcode.cn id=846 lang=golang
 *
 * [846] 一手顺子
 */

// 排序+思维
// 因为题目要求连续，每次从剩余的数中的最小值开始，遍历groupSize个，对遍历到的每个值的个数减一，遇到不够的即为false

package main

import (
	"sort"
)

// @lc code=start
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Slice(hand, func(i, j int) bool { return hand[i] < hand[j] })
	mp := map[int]int{}
	for i := range hand {
		mp[hand[i]]++
	}

	for _, num := range hand {
		if mp[num] == 0 {
			continue
		}

		k := num
		size := groupSize
		for size > 0 {
			size--
			mp[k]--
			if mp[k] < 0 {
				return false
			}
			k++
		}
	}

	return true
}

// @lc code=end

// func main() {
// 	fmt.Println(isNStraightHand([]int{1, 2, 3, 1, 2, 3, 6, 7, 8, 9, 7, 8}, 3))
// }
