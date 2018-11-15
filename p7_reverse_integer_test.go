package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

import "strconv"

var MAX = (1 << 31) - 1
var MIN = -(1 << 31)

// use a stack
// store the sign differently
// enqueue the digiti sequentially and then take them from the stack
// heading "0" should be skipped
func reverse(x int) int {
	positive := true
	if x < 0 {
		positive = false
		x = -x
	}

	stack := make([]int, 0)
	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		stack = append(stack, int(s[i]-'0'))
	}

	r := 0
	head := true

	for p, multipler := 0, 1; p < len(stack); p++ {
		// skip heading "0"
		if head && stack[p] == 0 {
			continue
		}

		head = false
		r = (r + stack[p]*multipler)

		multipler = multipler * 10
	}
	if !positive {
		r = -r
	}

	if r <= MIN || r >= MAX {
		return 0
	}
	return r
}

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(123, reverse(321))
	assert.Equal(0, reverse(0))
	assert.Equal(-123, reverse(-321))
	assert.Equal(0, reverse(1534236469))
}
