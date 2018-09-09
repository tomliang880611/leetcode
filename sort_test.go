package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{6, 7}

	actual := findMedianSortedArrays(nums1, nums2)
	if actual != 4 {
		t.Errorf("expected %d, got %f", 4, actual)
	}
}

func TestMergeSort(t *testing.T) {
	var rawArr = []int{12, 11, 13, 5, 6, 7}
	var expected = []int{5, 6, 7, 11, 12, 13}

	mergeSort(rawArr, 0, len(rawArr)-1)
	for i, v := range rawArr {
		if expected[i] != v {
			t.Errorf("expected %d, got %d", expected[i], v)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	var testData = []struct {
		array []int
		b     *testing.B
	}{
		{[]int{12, 11, 13, 5, 6, 7}, b},
		{[]int{12, 11, 13, 5, 6, 7, 100}, b},
		{[]int{12, 11, 13, 5, 6, 7, 12, 1}, b},
		{[]int{12, 11, 13, 5, 6, 7, 200, 1, 12}, b},
		{[]int{12, 11, 13, 5, 6, 7, 2, 1, 111, 100}, b},
	}
	for _, v := range testData {
		func(arr []int, b *testing.B) {
			for i := 0; i < b.N; i++ {
				mergeSort(arr, 0, len(arr)-1)
			}
		}(v.array, v.b)
	}
}

func TestQuickSort(t *testing.T) {
	data := []int{2, 1, 10, 333, 8, 9, 10}

	quickSort(data, 0, len(data)-1)
	fmt.Printf("sorted array %v", data)

	t.Error()
}
