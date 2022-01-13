/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// 265/265 cases passed (4 ms)
// Your runtime beats 37.35 % of golang submissions
// Your memory usage beats 65.57 % of golang submissions (2.5 MB)

// 因为是寻找  下一个  更大的排列，所以我们只需要找可以让排列大一点点的改动
// 倒序寻找递增的第一个下降点，那么这个点就是要被换掉的点
// 换成啥呢，换成倒序遍历过程中 恰好比他大一点点的值
// 以该点为中心，前面的数不动，后面数正序排列即可（因为后面的数原来是倒序，所以reverse一下就行）

// 举例
// 例如1, 3, 3, 5, 4, 2，那么倒序找到3，这个数需要被换掉
// 3换成遍历过程中的4
// 将然后将1,3,4,5,3,2的5,3,2 reverse 变成 1,3,4,2,3,5

package main

import "fmt"

// @lc code=start
func nextPermutation(nums []int) {

	p := len(nums) - 2
	for ; p >= 0 && nums[p] >= nums[p+1]; p-- {
	}

	if p >= 0 {
		q := len(nums) - 1
		for ; q > p && nums[q] <= nums[p]; q-- {
		}

		nums[p], nums[q] = nums[q], nums[p]
	}

	for i, j := p+1, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return
}

// @lc code=end

func main() {
	// nums := []int{1, 3, 3, 5, 4, 2}
	nums := []int{5}
	nextPermutation(nums)
	fmt.Println(nums)
}
