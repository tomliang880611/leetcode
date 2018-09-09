package leetcode

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	var haystack, needle string

	// case 1
	haystack = "hello"
	needle = "ll"

	if result := strStr(haystack, needle); result != 2 {
		t.Errorf("expecting 2 got, %v", result)
	}

	// case 2
	haystack = "aaaaa"
	needle = "bba"
	if result := strStr(haystack, needle); result != -1 {
		t.Errorf("expecting -1 got, %v", result)
	}

	haystack = "mississippi"
	needle = "sippia"
	if result := strStr(haystack, needle); result != -1 {
		t.Errorf("expecting -1 got, %v", result)
	}
}
