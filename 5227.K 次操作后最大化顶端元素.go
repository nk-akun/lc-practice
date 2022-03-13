package main

import "fmt"

func maximumTop(nums []int, k int) int {
	ans := -1
	if len(nums) > k-1 {
		for i := 0; i < k-1; i++ {
			if nums[i] > ans {
				ans = nums[i]
			}
		}
		if len(nums) > k && ans < nums[k] {
			ans = nums[k]
		}
	} else {
		flag := true
		maxNum := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				flag = false
			}
			if nums[i] > maxNum {
				maxNum = nums[i]
			}
		}
		if flag {
			k -= len(nums)
			if (k & 1) == 1 {
				ans = nums[0]
			}
		} else {
			ans = maxNum
		}
	}

	return ans
}

func main() {
	fmt.Println(maximumTop([]int{31, 15, 92, 84, 19, 92, 55}, 4))
}

// [31,15,92,84,19,92,55]
// 4
