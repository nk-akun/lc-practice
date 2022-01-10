package main

import "fmt"

func minSwaps(nums []int) int {
	nums1 := 0
	for _, n := range nums {
		if n == 1 {
			nums1++
		}
	}

	if nums1 == len(nums) {
		return 0
	}

	l, r := 0, 0
	maxNum := 0
	sum := 0
	for r < nums1 {
		if nums[r] == 1 {
			sum++
		}
		r++
	}
	maxNum = sum

	r = nums1 - 1
	for l < len(nums)-1 {
		if nums[l] == 1 {
			sum--
		}
		l++
		r = (r + 1) % len(nums)
		if nums[r] == 1 {
			sum++
		}
		if sum > maxNum {
			maxNum = sum
		}
		// fmt.Println(l, r, sum)
	}

	return nums1 - maxNum
}

func main() {
	fmt.Println(minSwaps([]int{1, 1, 1, 0, 0, 0, 0}))
}
