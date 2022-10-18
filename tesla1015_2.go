package main

import "fmt"

func Solution(S string, C []int) int {
	ans := 0
	con := []int{}
	for i := range S {
		if i == 0 {
			con = append(con, i)
		} else {
			if S[i] == S[con[len(con)-1]] {
				var minP int
				if C[i] > C[con[len(con)-1]] {
					minP = con[len(con)-1]
					con[len(con)-1] = i
				} else {
					minP = i
				}
				ans += C[minP]
			} else {
				con = append(con, i)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(Solution("abccbd", []int{0, 1, 2, 3, 4, 5}))
	fmt.Println(Solution("aabbcc", []int{1, 2, 1, 2, 1, 2}))
	fmt.Println(Solution("aaaa", []int{3, 4, 5, 6}))
	fmt.Println(Solution("ababa", []int{1, 2, 1, 2, 1}))
	fmt.Println(Solution("abaaba", []int{1, 2, 2, 2, 1, 2}))
	fmt.Println(Solution("aaa", []int{1, 2, 3}))
}
