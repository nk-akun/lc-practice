package main

func singleNumber(nums []int) int {
	var result int32 = 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range nums {
			count += num >> i & 1
		}
		if count%3 == 1 {
			result |= (1 << i)
		}
	}

	return int(result)
}

// func main() {

// 	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))

// }
