/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

/*
dfs深搜即可

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/

package main

import "fmt"

// @lc code=start

/*  方法1
var ans int

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	ans = 0x3f3f3f3f
	if (n & 1) == 1 {
		dfs(n+1, 1)
		dfs(n-1, 1)
	} else {
		dfs(n/2, 1)
	}
	return ans
}

func dfs(n int, step int) {
	if step >= ans {
		return
	}
	if n == 1 {
		if step < ans {
			ans = step
		}
		return
	}

	if (n & 1) == 1 {
		dfs(n+1, step+1)
		dfs(n-1, step+1)
	} else {
		dfs(n/2, step+1)
	}
	return
}
*/

// 方法二,如果二进制倒数第二位是1，当然选择+1，这样可以消掉好几个1；否则就-1
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	ans := 0
	for n != 1 {
		if (n & 1) == 1 {
			if ((n>>1)&1) == 1 && n > 3 {
				n++
			} else {
				n--
			}
		} else {
			n >>= 1
		}
		ans++
	}
	return ans
}

// @lc code=end

func main() {
	for i := 1<<31 - 1000; i <= 1<<31-1; i++ {
		fmt.Println(i, integerReplacement(i))
	}
}
