package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// this method has to be changed to reduce duplicate
// traversal of scanned elments
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var tmp = make(map[string][]int)
	// sort candidates in ascedant order
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})

	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}

		if candidates[i] == target {
			result = append(result, []int{candidates[i]})
			continue
		}

		val := candidates[i]
		// deal with sub match if not handled before
		reminder := combinationSum(candidates, target-val)
		for _, sub := range reminder {
			r := append(sub, val)
			// sort the result in asceding order so that the key
			// generated is stable
			sort.Slice(r, func(i, j int) bool {
				return r[i] < r[j]
			})
			// calculate unique key for a valid combination
			var key strings.Builder
			for _, v := range r {
				_, _ = key.WriteString(fmt.Sprintf("%v-", v))
			}
			tmp[key.String()] = r
		}
	}

	for _, r := range tmp {
		result = append(result, r)
	}
	return result
}

func TestCombinationSum(t *testing.T) {
	assert := assert.New(t)
	var candidate []int
	var target int
	var e [][]int
	candidate = []int{2, 3, 6, 7}
	target = 7
	e = [][]int{
		{7},
		{2, 2, 3},
	}

	// assert.Equal(e, combinationSum(candidate, target))

	candidate = []int{2, 3, 5}
	target = 8
	e = [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	// assert.Equal(e, combinationSum(candidate, target))

	candidate = []int{3, 5, 7}
	target = 10
	e = [][]int{
		{3, 7},
		{5, 5},
	}
	assert.Equal(e, combinationSum(candidate, target))

	// candidate = []int{3, 5, 7}
	// target = 5
	// e = [][]int{
	// 	{5},
	// }
	// assert.Equal(e, combinationSum(candidate, target))
}
