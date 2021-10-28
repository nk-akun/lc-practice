package main

func isValid(s string) bool {
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, len(s))
	top := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack[top] = s[i]
			top++
		} else if top > 0 {
			if stack[top-1] != mp[s[i]] {
				return false
			}
			top--
		} else {
			return false
		}
	}

	return top == 0
}

func main() {
	isValid("()[]{}")
}
