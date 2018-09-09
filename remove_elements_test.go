package leetcode

import "testing"

func TestRemoveElements(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	l := removeElement(arr, 3)
	if l != 4 {
		t.Errorf("failing to remove, expecting %v, got %v, the array is %v", 4, l, arr)
	}

	arr = []int{3, 2, 2, 3}
	l = removeElement(arr, 3)
	if l != 2 {
		t.Errorf("failing to remove, expecting %v, got %v, the array is %v", 2, l, arr)
	}
}
