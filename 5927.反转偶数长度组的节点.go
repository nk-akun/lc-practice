package main

// 这代码写的可太丑了

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	p := head
	i := 2
	for true {
		start, flag := judge(p, i)
		if flag {
			p = solve(start, i)
			// fmt.Println(p.Val)
		} else {
			p = start
			if start == nil {
				break
			}
		}
		if p == nil {
			break
		}
		i++
	}
	return head
}

func solve(p *ListNode, num int) *ListNode {
	if p.Next == nil {
		return nil
	}
	start := p
	q := p.Next
	for num > 0 && q != nil {
		np := q.Next
		q.Next = p
		p = q
		q = np
		num--
	}
	start.Next.Next = q
	ans := start.Next
	start.Next = p
	return ans
}

func judge(p *ListNode, num int) (*ListNode, bool) {
	start := p
	i := 0
	for i < num {
		p = p.Next
		if p == nil {
			if (i & 1) == 1 {
				return nil, false
			}
			return start, true
		}
		i++
	}
	// fmt.Println(i, num, q.Val)
	if (num & 1) == 1 {
		return p, false
	}
	return start, true
}

func main() {

}
