/*
 * @lc app=leetcode.cn id=458 lang=golang
 *
 * [458] 可怜的小猪
 */

// 依然用进制的方式解决此问题，这里引入轮次的概念，可尝试的轮次为k，那么需要的小猪的数量为
// 桶的数量用K+1进制表示时的长度，喂猪时对于当前桶的编号的k+1进制表示，第x位为i时，就将该
// 桶饮料在第i轮喂给第x头猪，由于只有一桶饮料有毒，这样k轮之后就可检测出有毒的那一桶

package main

import (
	"fmt"
	"math"
)

// @lc code=start
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	tryCount := minutesToTest / minutesToDie

	k := tryCount + 1
	b := float64(buckets)

	return int(math.Ceil(math.Log(b) / math.Log(float64(k))))
}

// @lc code=end

func main() {
	fmt.Println(poorPigs(4, 1, 2))
}
