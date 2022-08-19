package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	ans := 0
	cntMap := make(map[int]int, 0)
	cntMap[a[n-1]]++
	for j := n - 2; j > 0; j-- {
		for i := j - 1; i >= 0; i-- {
			ans += cntMap[3*a[j]-a[i]]
		}
		cntMap[a[j]]++
	}

	fmt.Println(ans)
}
