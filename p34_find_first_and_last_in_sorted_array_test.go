package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The idea was to identify the target first by using 1/2, the expand the
// the range

// searchRange was the method that return the range of the target
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	low, high := 0, len(nums)-1
	index := -1
	for low <= high {
		piv := (low + high) / 2
		if nums[piv] == target {
			index = piv
			break
		} else if target > nums[piv] {
			low = piv + 1
		} else {
			high = piv - 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	// expand and get the range
	first, last := index, index
	for first > 0 && nums[first-1] == target {
		first--
	}

	for last < len(nums)-1 && nums[last+1] == target {
		last++
	}

	return []int{first, last}

}

func TestSearchRange(t *testing.T) {
	assert := assert.New(t)
	var input []int
	var expect []int
	var target int

	// test case 1
	input = []int{1, 2, 3, 4, 4, 5}
	target = 4
	expect = []int{3, 4}

	r := searchRange(input, target)
	assert.Equal(expect, r)

	// test case 2
	input = []int{1, 2, 3, 4, 4, 5}
	target = -1
	expect = []int{-1, -1}

	r = searchRange(input, target)
	assert.Equal(expect, r)

	// test case 3
	input = []int{1}
	target = 1
	expect = []int{0, 0}

	r = searchRange(input, target)
	assert.Equal(expect, r)
}
