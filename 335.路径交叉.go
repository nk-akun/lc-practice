package main

func isSelfCrossing(distance []int) bool {
	if len(distance) <= 3 {
		return false
	}

	distance = append([]int{0}, distance...) // trick逻辑 {1,1,2,1,1}
	isJuaning := false
	for i := 2; i < len(distance); i++ {
		if !isJuaning && distance[i] <= distance[i-2] {
			// 开卷
			isJuaning = true
			continue
		}
		if isJuaning {
			// 情况1
			// 在真正内卷之前,是否撞上了之前的线段
			if i >= 5 && distance[i-2] > distance[i-4] {
				if distance[i-1] >= distance[i-3]-distance[i-5] && distance[i]+distance[i-4] >= distance[i-2] {
					return true
				}
			}

			// 情况2
			// 卷起来了
			if distance[i] >= distance[i-2] {
				return true
			}
		}
	}

	return false
}

// func main() {
// 	fmt.Println(isSelfCrossing([]int{0, 1, 2, 3, 4, 1, 3}))
// }
