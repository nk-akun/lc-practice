/*
 * @lc app=leetcode.cn id=1044 lang=golang
 *
 * [1044] 最长重复子串
 */

// 二分+字符串hash
// 最长的重复串的子串肯定也具备重复性，所以该题的重复子串的长度具备单调性
// 预处理出来字符串的hash值，然后二分长度，在map中记录子串的hash值出现的次数，大于等于2次即为重复

//  执行用时：204 ms, 在所有 Go 提交中击败了36.17%的用户
//  内存消耗：10.4 MB, 在所有 Go 提交中击败了19.15%的用户
//  通过测试用例：67 / 67

package main

// @lc code=start
func longestDupSubstring(s string) string {
	l, r, mid := 1, len(s)-1, -1
	hashS := calHash(s)

	ansl, ansr := -1, -1
	for l <= r {
		mid = (l + r) >> 1
		if start, end, ok := judge(s, hashS, mid); ok {
			ansl, ansr = start, end
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if ansl == -1 {
		return ""
	}
	return s[ansl : ansr+1]
}

func judge(s string, hashS []uint64, subLen int) (int, int, bool) {
	var k uint64 = 1000000007
	st := quickPow(k, subLen)

	hashValueMap := map[uint64]int{}
	l, r := 0, subLen-1
	for r < len(s) {
		var hashValue uint64
		if l == 0 {
			hashValue = hashS[r]
		} else {
			hashValue = hashS[r] - hashS[l-1]*st
		}
		if hashValueMap[hashValue] >= 1 {
			return l, r, true
		} else {
			hashValueMap[hashValue]++
		}
		l++
		r++
	}

	return -1, -1, false
}

func calHash(s string) []uint64 {
	var k uint64 = 1000000007
	hashS := make([]uint64, len(s))
	for i := range s {
		if i == 0 {
			hashS[i] = uint64(s[i] - 'a')
		} else {
			hashS[i] = hashS[i-1]*k + uint64(s[i]-'a')
		}
	}
	return hashS
}

func quickPow(k uint64, cnt int) uint64 {
	var result uint64 = 1
	for cnt > 0 {
		if cnt&1 == 1 {
			result *= k
		}
		k = k * k
		cnt >>= 1
	}
	return result
}

// @lc code=end

// func main() {
// 	fmt.Println(longestDupSubstring("acca"))
// }
