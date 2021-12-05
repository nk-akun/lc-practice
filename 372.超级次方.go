/*
 * @lc app=leetcode.cn id=372 lang=golang
 *
 * [372] 超级次方
 */

// 递归，x^123  =  x^120*x^3  =  (x^12)^10*(x^3)
// 执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.7 MB, 在所有 Go 提交中击败了58.73%的用户
// 通过测试用例：55 / 55
package main

import (
	"fmt"
)

// @lc code=start

const mod = 1337

func superPow(a int, b []int) int {
	return dfs(a, b)
}

func dfs(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	return quickPow(dfs(a, b[:len(b)-1]), 10) * quickPow(a, b[len(b)-1]) % mod
}

func quickPow(x int, num int) int {
	ans := 1
	for num > 0 {
		if num&1 == 1 {
			ans = ans * x % mod
		}
		num >>= 1
		x = x * x % mod
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(superPow(2, []int{3}))
}

// 将b数组不断除以2
/*
func superPow(a int, b []int) int {
	const mod = 1337
	ans := 1
	for !(len(b) == 0) {
		if b[len(b)-1]&1 == 1 {
			ans = ans * a % mod
		}
		b = div(b)
		a = a * a % mod
	}
	return ans % mod
}

func div(arr []int) []int {
	last := 0
	result := []int{}
	for i := 0; i < len(arr); i++ {
		now := last*10 + arr[i]
		result = append(result, now/2)
		last = now & 1
	}
	k := 0
	for k < len(result) && result[k] == 0 {
		k++
	}
	result = result[k:]
	return result
}

*/
