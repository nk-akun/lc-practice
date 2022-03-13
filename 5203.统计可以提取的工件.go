package main

import "fmt"

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	mp := make(map[[2]int]bool, len(dig))
	for i := range dig {
		mp[[2]int{dig[i][0], dig[i][1]}] = true
	}

	ans := 0
	for _, ar := range artifacts {
		flag := true
		for i := ar[0]; i <= ar[2]; i++ {
			for j := ar[1]; j <= ar[3]; j++ {
				if !flag {
					continue
				}
				flag = mp[[2]int{i, j}]
			}
		}
		if flag {
			// fmt.Println(ar)
			ans++
		}
	}
	return ans
}

func main() {
	// fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {1, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}}))
	// fmt.Println(digArtifacts(2, [][]int{{0, 0, 0, 0}, {1, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 1}}))
	fmt.Println(digArtifacts(2, [][]int{{3, 1, 4, 1}, {1, 1, 2, 2}, {1, 0, 2, 0}, {4, 3, 4, 4}, {0, 3, 1, 4}, {2, 3, 3, 4}}, [][]int{{0, 0}, {2, 1}, {2, 0}, {2, 3}, {4, 3}, {2, 4}, {4, 1}, {0, 2}, {4, 0}, {3, 1}, {1, 2}, {1, 3}, {3, 2}}))
}

// 5
// [[3,1,4,1],[1,1,2,2],[1,0,2,0],[4,3,4,4],[0,3,1,4],[2,3,3,4]]
// [[0,0],[2,1],[2,0],[2,3],[4,3],[2,4],[4,1],[0,2],[4,0],[3,1],[1,2],[1,3],[3,2]]
