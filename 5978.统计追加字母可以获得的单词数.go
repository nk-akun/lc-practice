package main

import "fmt"

func wordCount(startWords []string, targetWords []string) int {
	existMap := make(map[int]struct{})
	for i := range startWords {
		num := 0
		for j := range startWords[i] {
			num |= (1 << (startWords[i][j] - 'a'))
		}
		existMap[num] = struct{}{}
	}

	ans := 0
	for i := range targetWords {
		num := 0
		for j := range targetWords[i] {
			num |= (1 << (targetWords[i][j] - 'a'))
		}

		for j := range targetWords[i] {
			tmp := num ^ (1 << (targetWords[i][j] - 'a'))
			if _, ok := existMap[tmp]; ok {
				ans++
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(wordCount([]string{"g", "vf", "ylpuk", "nyf", "gdj", "j", "fyqzg", "sizec"}, []string{"r", "am", "jg", "umhjo", "fov", "lujy", "b", "uz", "y"}))
}

// ["g","vf","ylpuk","nyf","gdj","j","fyqzg","sizec"]
// ["r","am","jg","umhjo","fov","lujy","b","uz","y"]
