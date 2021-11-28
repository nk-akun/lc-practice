/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// 判断每个字符出现的次数。复杂度 O(n*26)
// 执行用时：4 ms, 在所有 Go 提交中击败了94.79%的用户
// 内存消耗：5 MB, 在所有 Go 提交中击败了97.73%的用户

package main

// @lc code=start
func findAnagrams(s string, p string) []int {
	numP := make([]int, 27)
	for _, c := range p {
		numP[c-'a']++
	}

	ans := []int{}
	numS := make([]int, 27)

	l, r := 0, 0
	for r < len(s) && r < len(p) {
		numS[s[r]-'a']++
		r++
	}

	r--
	for r < len(s) {
		// fmt.Println(l, r, numS, numP)
		if judge(numS, numP) {
			ans = append(ans, l)
		}

		numS[s[l]-'a']--
		l++
		r++
		if r < len(s) {
			numS[s[r]-'a']++
		}
	}

	return ans
}

func judge(s []int, p []int) bool {
	for i := 0; i < 26; i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}

// @lc code=end

// func main() {
// 	fmt.Println(findAnagrams("a", "f"))
// }

/*

// 可用于寻找s中p串的所有位置

func findAnagrams(s string, p string) []int {
	const (
		k   = 1000000007
		mod = 1000000009
	)
	hashS := make([]int64, len(s))
	hashS[0] = idx(s[0])
	for i := 1; i < len(s); i++ {
		hashS[i] = (hashS[i-1] + idx(s[i])*k) % mod
	}
	var hashP int64 = 0
	for i := 0; i < len(p); i++ {
		hashP = (hashP + idx(p[i])*k) % mod
	}

	ans := []int{}
	l, r := 0, len(p)-1
	for r < len(s) {
		var hashValue int64
		if l == 0 {
			hashValue = hashS[r]
		} else {
			hashValue = (hashS[r] - hashS[l-1]*quickPow(k, r-l+1) + mod) % mod
		}
		if hashValue == hashP {
			ans = append(ans, l)
		}
		l++
		r++
	}

	return ans
}

func idx(c byte) int64 {
	return int64(c - 'a')
}

func quickPow(x int64, p int) int64 {
	const mod = 1000000009
	var ans int64 = 1
	for p > 0 {
		if (p & 1) == 1 {
			ans = (ans * x) % mod
		}
		x = (x * x) % mod
		p >>= 1
	}
	return ans
}


*/
