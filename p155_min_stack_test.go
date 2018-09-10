package leetcode

import (
	"sort"
)

type node struct {
	val  int
	next *node
}
type MinStack struct {
	head   *node
	minArr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	n := &node{x, nil}
	n.next = this.head
	this.head = n

	this.minArr = append(this.minArr, x)
	sort.Ints(this.minArr)
}

func (this *MinStack) Pop() {
	n := this.head
	v := n.val
	if len(this.minArr) > 0 && this.minArr[0] == v {
		this.minArr = this.minArr[1:]
	}

	this.head = n.next
	n.next = nil
}

func (this *MinStack) Top() int {
	n := this.head
	return n.val
}

func (this *MinStack) GetMin() int {
	return this.minArr[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
