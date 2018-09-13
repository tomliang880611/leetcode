package leetcode

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// head of returning list
	var h *ListNode
	// current position
	var p *ListNode

	c1, c2 := l1, l2
	carrier := 0
	// continue if any list is not exhausted
	for c1 != nil || c2 != nil {
		var c1Val, c2Val int
		if c1 != nil {
			c1Val = c1.Val
		}
		if c2 != nil {
			c2Val = c2.Val
		}
		curSum := c1Val + c2Val + carrier
		carrier = curSum / 10
		reminder := curSum % 10

		c := &ListNode{Val: reminder}
		if h == nil {
			h, p = c, c
		} else {
			p.Next = c
		}
		p = c

		if c1 != nil {
			c1 = c1.Next
		}
		if c2 != nil {
			c2 = c2.Next
		}
	}

	// handle the last carrier
	if carrier > 0 {
		c := &ListNode{Val: carrier}
		p.Next = c
	}
	return h
}

func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := &ListNode{Val: 0}, &ListNode{Val: 0}

	r := addTwoNumbers(l1, l2)
	if r.Val != 0 {
		t.Errorf("expecting the first element being 0, got %v", r.Val)
	}
	if r.Next != nil {
		t.Errorf("Expecting single element, the next element being %v", r.Next)
	}
}
