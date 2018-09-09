package leetcode

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	i := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	e := [][]string{
		{"ate", "eat", "tea"},
		{"nat", "tan"},
		{"bat"},
	}
	r := groupAnagrams(i)
	if len(r) != len(e) {
		t.Errorf("expected %v arrays, got %v", len(e), len(r))
	}
}
