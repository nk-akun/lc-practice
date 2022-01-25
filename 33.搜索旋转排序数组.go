/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// 法1：直接通过二分法查找
// 取mid时，[l,mid]或[mid+1,r]至少一个是有序的
// 如果target在那个有序的区间内，直接将范围缩小至有序区间
// 否则将范围缩小至无序区间
// 相比与普通二分，多了一个判断命中区间的步骤

// 法2：先找到最大值top的位置，再判断target是在[0,top]中还是[top+1,length]中

// 195/195 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.4 MB)

package main

import "fmt"

// @lc code=start

// 法1
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if l == r {
			break
		}
		if nums[l] < nums[mid] {
			if target >= nums[l] && target <= nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			if target >= nums[mid+1] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return -1
}

// 法2
// func search(nums []int, target int) int {
// 	if nums[0] > nums[len(nums)-1] {
// 		top := findTop(nums)
// 		fmt.Println(top)
// 		if target >= nums[0] {
// 			return findTarget(nums[:top+1], target)
// 		} else {
// 			ans := findTarget(nums[top+1:], target)
// 			if ans != -1 {
// 				ans += (top + 1)
// 			}
// 			return ans
// 		}
// 	}
// 	return findTarget(nums, target)
// }

// func findTarget(nums []int, target int) int {
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if nums[mid] > target {
// 			r = mid - 1
// 		} else if nums[mid] == target {
// 			return mid
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }

// func findTop(nums []int) int {
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if nums[mid] >= nums[0] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return r
// }

// @lc code=end

func main() {
	fmt.Println(search([]int{9, 1, 3, 4, 5, 6, 7}, 10))
}
