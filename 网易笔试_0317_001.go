package main

import "fmt"

var ansNum int
var ansMap [][]int

func judge(tmap [][]int, g [][]int) bool {
	for i := 0; i < 3; i++ {
		vis := make([]bool, 4)
		for j := 0; j < 3; j++ {
			if vis[tmap[i][j]] {
				return false
			}
			vis[tmap[i][j]] = true
		}
	}
	for j := 0; j < 3; j++ {
		vis := make([]bool, 4)
		for i := 0; i < 3; i++ {
			if vis[tmap[i][j]] {
				return false
			}
			vis[tmap[i][j]] = true
		}
	}

	for i := 0; i < 3; i++ {
		vis := make([]bool, 4)
		for j := 0; j < 6; j += 2 {
			if vis[tmap[g[i][j]][g[i][j+1]]] {
				return false
			}
			vis[tmap[g[i][j]][g[i][j+1]]] = true
		}
	}

	return true
}

func dfs(x int, y int, smap []string, tmap [][]int, g [][]int) {
	if x < 2 && y == 3 {
		dfs(x+1, 0, smap, tmap, g)
		return
	}
	if x == 2 && y == 3 {
		if judge(tmap, g) {
			// copy(ansMap, tmap)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					ansMap[i][j] = tmap[i][j]
				}
			}
			ansNum++
		}
		return
	}

	if smap[x][y] == '*' {
		for i := 1; i <= 3; i++ {
			tmap[x][y] = i
			dfs(x, y+1, smap, tmap, g)
		}
	} else {
		tmap[x][y] = int(smap[x][y] - '0')
		dfs(x, y+1, smap, tmap, g)
	}
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		ansNum = 0
		ansMap = make([][]int, 3)
		for i := 0; i < 3; i++ {
			ansMap[i] = make([]int, 3)
		}
		smap := make([]string, 3)
		for i := range smap {
			fmt.Scanf("%s", &smap[i])
		}

		g := make([][]int, 3)
		for i := range g {
			g[i] = make([]int, 6)
			for j := range g[i] {
				fmt.Scanf("%d", &g[i][j])
			}
		}

		tmap := make([][]int, 3)
		for i := range tmap {
			tmap[i] = make([]int, 3)
		}
		dfs(0, 0, smap, tmap, g)

		if ansNum == 1 {
			fmt.Println("Unique")
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					fmt.Printf("%d", ansMap[i][j])
				}
				fmt.Println()
			}
		} else if ansNum > 1 {
			fmt.Println("Multiple")
		} else {
			fmt.Println("No")
		}
		t--
	}
}

/*

4
*2*
1*2
***
0 0 0 1 1 0
0 2 1 1 1 2
2 0 2 1 2 2
**3
***
***
0 0 1 0 1 1
0 1 0 2 1 2
2 0 2 1 2 2
**3
1**
**2
0 0 1 0 1 1
0 1 0 2 1 2
2 0 2 1 2 2
3*3
1**
**2
0 0 1 0 1 1
0 1 0 2 1 2
2 0 2 1 2 2

*/
