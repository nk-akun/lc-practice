package main

import (
	"strings"
)

func findWords(words []string) []string {
	keyb := [][]byte{{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}, {'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}, {'z', 'x', 'c', 'v', 'b', 'n', 'm'}}
	keyMap := map[byte]int{}
	for i := range keyb {
		for j := range keyb[i] {
			keyMap[keyb[i][j]] = i
		}
	}

	result := []string{}
	for i, word := range words {
		word = strings.ToLower(word)
		oneLine := true
		for i := 1; i < len(word); i++ {
			if keyMap[word[i]] != keyMap[word[i-1]] {
				oneLine = false
				break
			}
		}
		if oneLine {
			result = append(result, words[i])
		}
	}

	return result
}

// func main() {
// 	fmt.Println(findWords([]string{"adSdf", "sfSSSSSSSSSd"}))
// }

/*

执行结果：通过
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了42.11%的用户
通过测试用例：22 / 22

*/
