package main

import "fmt"

func Solution(S string) int {
	pos := make([]bool, len(S)+1)
	for i := 1; i <= len(S); i++ {
		pos[i] = true
	}

	ans := 0
	for i := range S {
		if S[i] == '<' {
			if !pos[i] {
				ans++
				pos[i+1] = false
			}
		} else if S[i] == '>' {
			if i == len(S)-1 {
				ans++
			}
		} else if S[i] == '^' || S[i] == 'v' {
			ans++
			pos[i+1] = false
		}
	}
	return ans
}

func main() {
	fmt.Println(Solution("><^v"))
	fmt.Println(Solution("<<^<v>>"))
	fmt.Println(Solution("><><"))
	fmt.Println(Solution("^v^v^v<"))
}
