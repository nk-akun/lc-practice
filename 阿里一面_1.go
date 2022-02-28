package main

// 交换单链表相邻节点
// 例如 1,2,3,4,5
// 结果 2,1,4,3,5

import "fmt"

type node struct {
	value int
	next  *node
}

func swap(phead *node) *node {
	var last *node
	p := phead
	q := p.next
	phead = q

	if q == nil {
		return p
	}

	for q != nil {
		// 交换
		p.next = q.next
		q.next = p
		if last != nil {
			last.next = q
		}
		q = p.next
		if q == nil {
			break
		}
		last = p

		// 后移
		p = q
		q = q.next
	}

	return phead
}

func main() {
	// list := []int{1, 2, 3, 4, 5}
	// list := []int{1, 2}
	list := []int{1}
	var head, last *node

	for i, x := range list {
		if i == 0 {
			head = &node{
				value: x,
				next:  nil,
			}
			last = head
		} else {
			now := &node{
				value: x,
				next:  nil,
			}
			last.next = now
			last = now
		}
	}

	p := head
	for p != nil {
		fmt.Print(p.value)
		p = p.next
	}
	fmt.Println()

	phead := swap(head)
	for phead != nil {
		fmt.Print(phead.value)
		phead = phead.next
	}
	fmt.Println()
}
