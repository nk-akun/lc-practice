package main

import (
	"math"
)

func solveC(my []int64, k int64) int64 {
	if my[0] != 0 {
		return 1
	}

	var result int64 = 0
	var i int
	for i = 1; i < len(my); i++ {
		bill := int64(math.Pow10(int(my[i-1])))
		needNum := int64(math.Pow10(int(my[i]-my[i-1]))) - 1
		if k >= needNum {
			result += bill * needNum
			k -= needNum
		} else {
			break
		}
	}

	bill := int64(math.Pow10(int(my[i-1])))
	needNum := k + 1 // 求凑不成的数，所以需要k+1
	result += bill * needNum

	return result
}

// func main() {
// 	input := bufio.NewReader(os.Stdin)

// 	var n int
// 	var t, k int64
// 	fmt.Fscan(input, &t)

// 	for t > 0 {
// 		t--
// 		fmt.Fscan(input, &n, &k)
// 		my := make([]int64, n)
// 		for i := 0; i < n; i++ {
// 			fmt.Fscan(input, &my[i])
// 		}

// 		result := solveC(my, k)
// 		fmt.Println(result)
// 	}
// }
