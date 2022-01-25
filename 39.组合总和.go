/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// 上楼梯走一步走两步问题
// 需要记录方案，所以需要把走到每个台阶的方案记录下来，走向下个台阶时需要把方案也转移过去

// 170/170 cases passed (4 ms)
// Your runtime beats 60.35 % of golang submissions
// Your memory usage beats 6.04 % of golang submissions (6.5 MB)

package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][][]int, target+1)
	for i := 0; i < target+1; i++ {
		ans[i] = make([][]int, 0)
	}

	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	for i := range candidates {
		if candidates[i] > target {
			break
		}
		ans[candidates[i]] = append(ans[candidates[i]], []int{candidates[i]})
	}

	for i := 1; i <= target; i++ {
		for _, k := range candidates {
			if i <= k || len(ans[i-k]) == 0 {
				continue
			}
			for j := range ans[i-k] {
				tmp := make([]int, len(ans[i-k][j]))
				copy(tmp, ans[i-k][j]) // 深拷贝
				tmp = append(tmp, k)
				if !judgeExist(ans[i], tmp) {
					ans[i] = append(ans[i], tmp)
				}
			}
		}
	}
	return ans[target]
}

func judgeExist(all [][]int, tmp []int) bool {
	sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
	for l := range all {
		if len(all[l]) != len(tmp) {
			continue
		}
		i := 0
		for ; i < len(tmp); i++ {
			if all[l][i] != tmp[i] {
				break
			}
		}
		if i == len(tmp) {
			return true
		}
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}
