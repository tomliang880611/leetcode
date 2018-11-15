package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	max := 0

	low := 0
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if pos, ok := seen[cur]; ok {
			// only the char seen after the last one is valid
			if pos+1 > low {
				low = pos + 1
			}
		} else if diff := i - low + 1; diff > max {
			max = diff
		}
		seen[cur] = i
	}

	return max
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var input = "abcabcbb"
	var expect = 3
	assert := assert.New(t)
	assert.Equal(expect, lengthOfLongestSubstring(input))

	input = "bbbbb"
	expect = 1
	assert.Equal(expect, lengthOfLongestSubstring(input))

	input = "pwwkew"
	expect = 3
	assert.Equal(expect, lengthOfLongestSubstring(input))

	input = "au"
	expect = 2
	assert.Equal(expect, lengthOfLongestSubstring(input))

	input = " "
	expect = 1
	assert.Equal(expect, lengthOfLongestSubstring(input))

	input = ""
	expect = 0
	assert.Equal(expect, lengthOfLongestSubstring(input))

	input = "abba"
	expect = 2
	assert.Equal(expect, lengthOfLongestSubstring(input))
}
