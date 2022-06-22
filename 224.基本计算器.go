/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

package main

// 42/42 cases passed (4 ms)
// Your runtime beats 53.28 % of golang submissions
// Your memory usage beats 13.13 % of golang submissions (6.3 MB)

// 双栈 数字栈 符号栈
// 添加一个数字，符合栈push一个#
// 添加一个符号，数字栈push一个0
// 保证两个栈长度一致

// 遇到右括号时，开始pop栈，直到在符号栈遇到左括号为止，将计算结果push进数字栈，同时符号栈push一个‘#’
// 例如
// [0 0 12 0 3 0 1 0 1 0 1 0 2]  数字栈
// (-#+#-#-#-#+#                 符号栈

// 为保证最终数字栈就剩一个数字，给s加上左右括号，s = "(" + s + ")"

import "fmt"

// @lc code=start
func calculate(s string) int {
	s = "(" + s + ")"
	stNum := []int{}
	stOp := []byte{}
	var num int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			num, i = findNum(s, i)
			stNum = append(stNum, num)
			stOp = append(stOp, '#')
		} else if s[i] == ')' {
			result := 0
			pos := len(stNum) - 1
			for stOp[pos] != '(' {
				if stOp[pos] == '(' {
					break
				}
				if stOp[pos] != '#' {
					pos--
					continue
				}
				if pos > 0 && stOp[pos-1] == '-' {
					result += -stNum[pos]
				} else {
					result += stNum[pos]
				}
				pos--
			}
			stNum = stNum[:pos]
			stOp = stOp[:pos]

			stNum = append(stNum, result)
			stOp = append(stOp, '#')
		} else {
			stNum = append(stNum, 0)
			stOp = append(stOp, s[i])
		}
	}

	return stNum[0]
}

func findNum(s string, pos int) (int, int) {
	for s[pos] < '0' || s[pos] > '9' {
		pos++
	}
	result := 0
	for pos < len(s) && (s[pos] >= '0' && s[pos] <= '9') {
		result = result*10 + int(s[pos]-'0')
		pos++
	}
	return result, pos - 1
}

// @lc code=end

func main() {
	fmt.Println(calculate("1"))
	fmt.Println(calculate("1+1"))
	fmt.Println(calculate(" 2-1 + 2"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("-12 + 3-1-1-1+2"))
	fmt.Println(calculate("-(1-1)"))
}
