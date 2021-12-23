/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 */

// KMP裸题，先求模式串也就是b的next数组，然后假设a很长很长，不断匹配就可以了，匹配的时候i%len(a)就能实现a的循环了
// KMP算法参考链接：https://www.acwing.com/blog/content/2347/

//  57/57 cases passed (0 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 70.67 % of golang submissions (2.5 MB)

package main

// @lc code=start
// func repeatedStringMatch(a string, b string) int {
// 	i, j := 0, 0
// 	var maxl int
// 	if len(a) > len(b) {
// 		maxl = 2 * len(a)
// 	} else {
// 		maxl = 2 * len(b)
// 	}

// 	next := getNext(b)
// 	for ; i <= maxl+1; i++ {
// 		for j != -1 && a[i%len(a)] != b[j] {
// 			j = next[j]
// 		}
// 		if j == -1 || a[i%len(a)] == b[j] {
// 			j++
// 		}
// 		if j == len(b) {
// 			i++
// 			break
// 		}
// 	}
// 	if i > maxl+1 {
// 		return -1
// 	}

// 	return (i + len(a) - 1) / len(a)
// }

// func getNext(s string) []int {
// 	next := make([]int, len(s)+1)

// 	i, j := 0, -1
// 	next[0] = -1
// 	for i < len(s) {
// 		if j == -1 || s[i] == s[j] {
// 			i++
// 			j++
// 			next[i] = j
// 		} else {
// 			j = next[j]
// 		}
// 	}
// 	return next
// }

// 通过计算哈希值的方式寻找b串，计算b的哈希值，然后向上面kmp一样循环计算a的哈希值，在其中寻找b
// 关键词：用uint64，st值的提前计算，哈希值的计算方式：h = h*k+idx(c)

// 57/57 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 65.33 % of golang submissions (4.7 MB)
func repeatedStringMatch(a string, b string) int {
	var (
		k  uint64 = 1000000007
		st uint64 = 1
	)

	var maxl int
	if len(a) > len(b) {
		maxl = 2 * len(a)
	} else {
		maxl = 2 * len(b)
	}

	// 计算区间哈希值用
	for i := 0; i < len(b); i++ {
		st *= k
	}

	// 计算b的哈希值
	var hashB uint64 = 0
	for i := range b {
		hashB = hashB*k + idx(b[i])
	}

	// 一边计算一边匹配
	hashS := []uint64{}
	for i := 0; i < maxl+1; i++ {
		if i == 0 {
			hashS = append(hashS, idx(a[0]))
		} else {
			hashS = append(hashS, hashS[i-1]*k+idx(a[i%len(a)]))
		}
		if len(hashS) < len(b) {
			continue
		}
		var hashValue uint64
		if len(hashS) == len(b) {
			hashValue = hashS[i]
		} else {
			hashValue = hashS[i] - hashS[i-len(b)]*st
		}
		if hashValue == hashB {
			i++
			return (i + len(a) - 1) / len(a)
		}
	}
	return -1
}

func idx(c byte) uint64 {
	return uint64(c - 'a')
}

// @lc code=end

// func main() {
// 	fmt.Println(repeatedStringMatch("abc", "wxyz"))
// }
