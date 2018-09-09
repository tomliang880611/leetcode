package leetcode

import (
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	var i []int
	var o, e int

	i = []int{1, 3, 5, 6}
	o = 5
	e = 2
	if r := searchInsert(i, o); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	i = []int{1, 3, 5, 6}
	o = 2
	e = 1

	if r := searchInsert(i, o); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}
	i = []int{1, 3, 5, 6}
	o = 7
	e = 4
	if r := searchInsert(i, o); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	i = []int{1, 3, 5, 6}
	o = 0
	e = 0
	if r := searchInsert(i, o); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	i = []int{1}
	o = 1
	e = 0
	if r := searchInsert(i, o); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	i = []int{1}
	o = 2
	e = 1
	if r := searchInsert(i, o); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}
}
