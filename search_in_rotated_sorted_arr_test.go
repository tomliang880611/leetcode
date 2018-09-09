package leetcode

import (
	"testing"
)

func TestSearchInRotatedSortedArr(t *testing.T) {
	var nums []int
	var target int
	var expectation int

	// case 1
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	expectation = 4
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	// case 2 no found
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	expectation = -1
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{1, 3}
	target = 0
	expectation = -1
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{1}
	target = 1
	expectation = 0
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{1, 3}
	target = 3
	expectation = 1
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 2
	expectation = 6
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	target = 1
	expectation = 6
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{5, 1, 2, 3, 4}
	target = 1
	expectation = 1
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{1, 2, 3, 4, 5, 6}
	target = 4
	expectation = 3
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 2
	expectation = 6
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 6
	expectation = 2
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}
}

func TestSearchInRotatedSortedArr2(t *testing.T) {
	var nums []int
	var target int
	var expectation int

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 6
	expectation = 2
	if r := search(nums, target); r != expectation {
		t.Errorf("expecting %v, got %v", expectation, r)
	}

}
