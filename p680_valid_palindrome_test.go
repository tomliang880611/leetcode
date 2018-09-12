package leetcode

import (
	"testing"
)

func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] != s[high] {
			return doValid(low+1, high, s) || doValid(low, high-1, s)
		}
		high--
		low++
	}

	return true
}

func doValid(low, high int, s string) bool {
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--

	}
	return true
}

func TestValidPalindrome(t *testing.T) {
	var input string
	var e bool

	// test case one
	input = "abca"
	e = true

	if r := validPalindrome(input); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	// test case two
	input = "aba"
	e = true

	if r := validPalindrome(input); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}
	// test case three
	input = "ab"
	e = true

	if r := validPalindrome(input); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	// test case four
	input = "abc"
	e = false

	if r := validPalindrome(input); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	input = "deeee"
	e = true

	if r := validPalindrome(input); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}

	input = "eeee"
	e = true

	if r := validPalindrome(input); r != e {
		t.Errorf("expecting %v, got %v", e, r)
	}
}
