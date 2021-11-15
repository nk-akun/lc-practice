/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// 使用Trie来解，节点记录以此节点结尾的前缀的值，插入时更新每个前缀的值

package main

// @lc code=start
type MapSum struct {
	mp   map[string]int
	root *Node
}

type Node struct {
	value int
	next  []*Node
}

func NewNode() *Node {
	return &Node{
		value: 0,
		next:  make([]*Node, 26),
	}
}

func Constructor() MapSum {
	return MapSum{
		mp:   map[string]int{},
		root: NewNode(),
	}
}

func (this *MapSum) Insert(key string, val int) {
	orgin := val
	if v, exist := this.mp[key]; exist {
		val -= v
	}
	this.mp[key] = orgin

	p := this.root
	for _, c := range key {
		pos := c - 'a'
		if p.next[pos] == nil {
			p.next[pos] = NewNode()
		}
		p.next[pos].value += val
		p = p.next[pos]
	}
	return
}

func (this *MapSum) Sum(prefix string) int {
	p := this.root
	for _, c := range prefix {
		pos := c - 'a'
		if p.next[pos] == nil {
			return 0
		}
		p = p.next[pos]
	}
	return p.value
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

// func main() {
// 	mapSum := Constructor()
// 	m := &mapSum

// 	m.Insert("apple", 3)
// 	fmt.Println(m.Sum("ap"))
// 	m.Insert("apple", 2)
// 	fmt.Println(m.Sum("ap"))
// }
