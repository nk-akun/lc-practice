package main

// lowbit将两个数的不同位找到,再对全体数组分类

func singleNumber(nums []int) []int {

	var mask int = 0
	for _, num := range nums {
		mask ^= num
	}

	// lowbit
	mask = mask & (-mask)

	result := make([]int, 2)
	for _, num := range nums {
		if (num & mask) == 0 {
			result[0] = result[0] ^ num
		} else {
			result[1] = result[1] ^ num
		}
	}

	return result
}

// func main() {
// 	fmt.Println(singleNumber([]int{1, 2, 1, 2, 3, 4}))
// }

/*
32 / 32 个通过测试用例
状态：通过
执行用时: 4 ms
内存消耗: 3.9 MB
*/
