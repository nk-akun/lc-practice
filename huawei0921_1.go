package main

import "fmt"

var n int
var total int
var a []int

func main() {

	var sum int = 0
	fmt.Scan(&n)
	a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	fmt.Scan(&total)

	if sum <= total {
		fmt.Println(-1)
		return
	}

	l, r, mid := 1, 1000000, 0
	for l <= r {
		mid = (l + r) >> 1
		count := 0
		for i := 1; i <= n; i++ {
			if a[i] > mid {
				count += mid
			} else {
				count += a[i]
			}
		}
		if count > total {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	fmt.Println(r)
}
