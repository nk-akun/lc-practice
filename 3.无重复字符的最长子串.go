package main

// 尺取直接做

func lengthOfLongestSubstring(s string) int {
	exist := map[byte]struct{}{}

	l, r, maxLen := 0, 0, 0
	for r < len(s) {
		if _, ok := exist[s[r]]; !ok {
			exist[s[r]] = struct{}{}
			r++
			maxLen = max(maxLen, r-l)
		} else {
			delete(exist, s[l])
			l++
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcdbcdef"))
// }
