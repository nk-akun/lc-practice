package main

import "fmt"

func judge(state int, aliceArrows []int, numArrows int) (int, []int) {
	result := 0
	bob := []int{}
	for i := 0; i < 12; i++ {
		if state&(1<<i) > 0 {
			result += i
			numArrows -= (aliceArrows[i] + 1)
			if numArrows < 0 {
				return 0, []int{}
			}
			bob = append(bob, aliceArrows[i]+1)
			continue
		}
		bob = append(bob, 0)
	}
	if numArrows > 0 {
		bob[11] += numArrows
	}
	return result, bob
}

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	ans := 0
	bob := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < (1<<12)-1; i++ {
		tmp1, tmp2 := judge(i, aliceArrows, numArrows)
		if tmp1 > ans {
			ans = tmp1
			bob = tmp2
		}
	}
	// fmt.Println(ans)
	return bob
}

func main() {
	fmt.Println(maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}))
	fmt.Println(maximumBobPoints(0, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
