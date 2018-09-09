package leetcode

import (
	"testing"
)

func TestRotateImage(t *testing.T) {
	var matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix)
	var e = [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	for i := 0; i < len(e); i++ {
		for j := 0; j < len(e); j++ {
			if matrix[i][j] != e[i][j] {
				t.Errorf("at [%v, %v], %v is expected, got %v.", i, j, e[i][j], matrix[i][j])
			}
		}
	}
}
