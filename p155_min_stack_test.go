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

func (s *MinStack) Push(x int) {
	n := &node{x, nil}
	n.next = s.head
	s.head = n

	s.minArr = append(s.minArr, x)
	sort.Ints(s.minArr)
}

func (s *MinStack) Pop() {
	n := s.head
	v := n.val
	if len(s.minArr) > 0 && s.minArr[0] == v {
		s.minArr = s.minArr[1:]
	}

	s.head = n.next
	n.next = nil
}

func (s *MinStack) Top() int {
	n := s.head
	return n.val
}

func (s *MinStack) GetMin() int {
	return s.minArr[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
