package main

// 直接异或

func singleNumber1(nums []int) int {
	var result int = 0

	for _, num := range nums {
		result ^= num
	}

	return result
}

// func main() {
// 	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
// }
