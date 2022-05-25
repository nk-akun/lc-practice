/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

package main

// 92/92 cases passed (40 ms)
// Your runtime beats 47.35 % of golang submissions
// Your memory usage beats 5.25 % of golang submissions (7.9 MB)

// 前缀和  sum[i]-sum[j]=k ==> sum[j]=sum[i]-k
// 所以对于当前前缀和 sum[i],计算count[sum[i]-k]的个数，即为以i为右边界的组成k的区间数

// @lc code=start
func subarraySum(nums []int, k int) int {
	ans := 0

	countMap := make(map[int]int)
	sum := 0
	countMap[sum]++
	for i := range nums {
		sum += nums[i]
		ans += countMap[sum-k]
		countMap[sum]++
	}

	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
// 	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
// 	fmt.Println(subarraySum([]int{1, 2, 3, 4, 4, 3, 2, 1, 1, 2, 3, 4, 5}, 9))
// 	fmt.Println(subarraySum([]int{1}, 1))
// }
