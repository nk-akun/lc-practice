package main

import "fmt"

func judge(nums []int, l, r int) int {
	// fmt.Println(len(nums), l, r)
	if l == 0 || r >= len(nums)-1 {
		return 0
	}
	last := nums[l-1]
	next := nums[r+1]
	if last < nums[l] && next < nums[l] || last > nums[l] && next > nums[l] {
		return 1
	}
	return 0
}

func countHillValley(nums []int) int {
	ans := 0
	for l, r := 0, 0; l < len(nums)-1; l = r + 1 {
		for r = l + 1; r < len(nums); r++ {
			if nums[r] != nums[r-1] {
				r--
				break
			}
		}
		ans += judge(nums, l, r)
	}
	return ans
}

func main() {
	fmt.Println(countHillValley([]int{57, 57, 57, 57, 57, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 85, 85, 85, 86, 86, 86}))
}

// [57,57,57,57,57,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,85,85,85,86,86,86]
