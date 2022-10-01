package main

import "fmt"

var dp []int
var height []int

const mod = 1000000007

func main() {
	var n, p, k int
	fmt.Scan(&n)
	fmt.Scan(&p)
	fmt.Scan(&k)

	dp = make([]int, 100050)
	height = make([]int, 100050)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		for j := 1; j <= num; j++ {
			fmt.Scan(&height[j])
		}
		for j := 0; j <= num; j++ {
			dp[j] = 0
		}
		dp[0] = 1
		for j := 1; j <= num; j++ {
			sum := height[j]
			for l := j - 1; l >= j-k && l >= 0; l-- {
				if sum > p {
					break
				}
				sum = sum + height[l]
				dp[j] = (dp[j] + dp[l]) % mod
			}
		}
		fmt.Println(dp[num])
	}

}
