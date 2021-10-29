package main

func reorderedPowerOf2(n int) bool {

	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	return dfs(nums, 0)
}

func dfs(nums []int, pos int) bool {
	if pos == len(nums) {
		return judge(nums)
	}
	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos]
		if dfs(nums, pos+1) {
			return true
		}
		nums[pos], nums[i] = nums[i], nums[pos]
	}
	return false
}

func judge(nums []int) bool {
	result := 0
	if nums[0] == 0 {
		return false
	}

	for _, x := range nums {
		result = result*10 + x
	}

	return result&(result-1) == 0
}

// func main() {
// 	fmt.Println(reorderedPowerOf2(987654321))
// }
